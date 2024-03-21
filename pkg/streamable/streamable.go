package streamable

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/chia-network/go-chia-libs/pkg/util"
)

const (
	// Name of the struct tag used to identify the streamable properties
	tagName = "streamable"

	// Bytes that indicate bool yes or no when serialized
	boolFalse uint8 = 0
	boolTrue  uint8 = 1
)

// Unmarshal unmarshals a streamable type based on struct tags
// Struct order is extremely important in this decoding. Ensure the order/types are identical
// on both sides of the stream
func Unmarshal(bytes []byte, v interface{}) error {
	tv := reflect.ValueOf(v)
	if tv.Kind() != reflect.Ptr || tv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	// Gets rid of the pointer
	tv = reflect.Indirect(tv)

	// Get the actual type
	t := tv.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("streamable can't unmarshal into non-struct type")
	}

	_, err := unmarshalStruct(bytes, t, tv)
	return err
}

func unmarshalStruct(bytes []byte, t reflect.Type, tv reflect.Value) ([]byte, error) {
	var err error

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		fieldValue := tv.Field(i)
		fieldType := fieldValue.Type()

		bytes, err = unmarshalField(bytes, fieldType, fieldValue, structField)
		if err != nil {
			return bytes, err
		}
	}

	return bytes, nil
}

// Struct field is used to parse out the streamable tag
// Not needed for anything else
// When recursively calling this on a wrapper type like mo.Option, pass the parent/wrapping StructField
func unmarshalField(bytes []byte, fieldType reflect.Type, fieldValue reflect.Value, structField reflect.StructField) ([]byte, error) {
	if _, tagPresent := structField.Tag.Lookup(tagName); !tagPresent {
		// Continuing because the tag isn't present
		return bytes, nil
	}

	var err error
	var newVal []byte

	// Optionals are handled with mo.Option
	// There will be one byte bool that indicates whether the field is present
	if strings.HasPrefix(fieldType.String(), "mo.Option[") {
		var presentFlag []byte
		presentFlag, bytes, err = util.ShiftNBytes(1, bytes)
		if err != nil {
			return bytes, err
		}
		if presentFlag[0] == boolFalse {
			return bytes, nil
		}

		// The unsafe.Pointer(..) stuff in here is to be able to set unexported fields of mo.Option
		// See https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields

		// First we set the present attr to true
		presentField := fieldValue.Field(0)
		presentFieldType := presentField.Type()
		reflect.NewAt(presentFieldType, unsafe.Pointer(presentField.UnsafeAddr())).Elem().SetBool(true)

		// Then, we can parse out the value of the field and set the value attr
		optionalField := fieldValue.Field(1)
		optionalType := optionalField.Type()
		writableField := reflect.NewAt(optionalType, unsafe.Pointer(optionalField.UnsafeAddr())).Elem()
		bytes, err = unmarshalField(bytes, optionalType, writableField, structField)
		if err != nil {
			return bytes, err
		}

		return bytes, nil
	}

	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()

		// Need to init the field to something non-nil before using it
		fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
		fieldValue = fieldValue.Elem()
	}

	switch kind := fieldType.Kind(); kind {
	case reflect.Uint8:
		newVal, bytes, err = util.ShiftNBytes(1, bytes)
		if err != nil {
			return bytes, err
		}
		if !fieldValue.CanSet() {
			return bytes, fmt.Errorf("field %s is not settable", fieldValue.String())
		}
		fieldValue.SetUint(uint64(util.BytesToUint8(newVal)))
	case reflect.Uint16:
		newVal, bytes, err = util.ShiftNBytes(2, bytes)
		if err != nil {
			return bytes, err
		}
		if !fieldValue.CanSet() {
			return bytes, fmt.Errorf("field %s is not settable", fieldValue.String())
		}
		newInt := util.BytesToUint16(newVal)
		fieldValue.SetUint(uint64(newInt))
	case reflect.Uint32:
		newVal, bytes, err = util.ShiftNBytes(4, bytes)
		if err != nil {
			return bytes, err
		}
		if !fieldValue.CanSet() {
			return bytes, fmt.Errorf("field %s is not settable", fieldValue.String())
		}
		newInt := util.BytesToUint32(newVal)
		fieldValue.SetUint(uint64(newInt))
	case reflect.Uint64:
		newVal, bytes, err = util.ShiftNBytes(8, bytes)
		if err != nil {
			return bytes, err
		}
		if !fieldValue.CanSet() {
			return bytes, fmt.Errorf("field %s is not settable", fieldValue.String())
		}
		newInt := util.BytesToUint64(newVal)
		fieldValue.SetUint(newInt)
	case reflect.Array:
		switch fieldValue.Type().Elem().Kind() {
		case reflect.Uint8:
			for i := 0; i < fieldValue.Len(); i++ {
				optionalField := fieldValue.Index(1)
				optionalType := optionalField.Type()
				bytes, err = unmarshalField(bytes, optionalType, fieldValue.Index(i), structField)
				if err != nil {
					return bytes, err
				}
			}
		default:
			return bytes, fmt.Errorf("unimplemented array type %s", fieldType.Elem().Kind())
		}
	case reflect.Slice:
		var length []byte
		length, bytes, err = util.ShiftNBytes(4, bytes)
		if err != nil {
			return bytes, err
		}
		numItems := binary.BigEndian.Uint32(length)

		sliceReflect := reflect.MakeSlice(fieldValue.Type(), 0, 0)
		for j := uint32(0); j < numItems; j++ {
			newValue := reflect.Indirect(reflect.New(fieldValue.Type().Elem()))
			bytes, err = unmarshalField(bytes, fieldType.Elem(), newValue, structField)
			if err != nil {
				return bytes, err
			}
			sliceReflect = reflect.Append(sliceReflect, newValue)
		}

		fieldValue.Set(sliceReflect)
	case reflect.String:
		// 4 byte size prefix, then []byte which can be converted to utf-8 string
		// Get 4 byte length prefix
		var length []byte
		length, bytes, err = util.ShiftNBytes(4, bytes)
		if err != nil {
			return bytes, err
		}
		numBytes := binary.BigEndian.Uint32(length)

		var strBytes []byte
		strBytes, bytes, err = util.ShiftNBytes(uint(numBytes), bytes)
		if err != nil {
			return bytes, err
		}
		fieldValue.SetString(string(strBytes))
	case reflect.Struct:
		bytes, err = unmarshalStruct(bytes, fieldType, fieldValue)
		if err != nil {
			return bytes, err
		}
	case reflect.Bool:
		var boolByte []byte
		boolByte, bytes, err = util.ShiftNBytes(1, bytes)
		if err != nil {
			return bytes, err
		}
		fieldValue.SetBool(boolByte[0] == boolTrue)
	default:
		return bytes, fmt.Errorf("unimplemented type %s", fieldValue.Kind())
	}

	return bytes, nil
}

// Marshal marshals the item into the streamable byte format
func Marshal(v interface{}) ([]byte, error) {
	// Doesn't matter if a pointer or not for marshalling, so
	// we just call this and let it deal with ptr or not ptr
	tv := reflect.Indirect(reflect.ValueOf(v))

	// Get the actual type
	t := tv.Type()

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("streamable can't marshal a non-struct type")
	}

	// This will become the final encoded data
	var finalBytes []byte

	return marshalStruct(finalBytes, t, tv)
}

func marshalStruct(finalBytes []byte, t reflect.Type, tv reflect.Value) ([]byte, error) {
	var err error

	// Iterate over all available fields in the type and encode to bytes
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		fieldValue := tv.Field(i)
		fieldType := fieldValue.Type()

		finalBytes, err = marshalField(finalBytes, fieldType, fieldValue, structField)
		if err != nil {
			return finalBytes, err
		}
	}

	return finalBytes, nil
}

func marshalField(finalBytes []byte, fieldType reflect.Type, fieldValue reflect.Value, structField reflect.StructField) ([]byte, error) {
	var err error

	var tagPresent bool
	if _, tagPresent = structField.Tag.Lookup(tagName); !tagPresent {
		// Continuing because the tag isn't present
		return finalBytes, nil
	}

	// Optionals are handled with mo.Option
	// If the value is not present, 0x00 and move on
	// otherwise 0x01 and encode the value
	if strings.HasPrefix(fieldType.String(), "mo.Option[") {
		isPresent := fieldValue.MethodByName("IsPresent").Call([]reflect.Value{})[0].Bool()

		if !isPresent {
			// Field is not present, insert `false` byte and continue
			finalBytes = append(finalBytes, boolFalse)
			return finalBytes, nil
		}

		finalBytes = append(finalBytes, boolTrue)

		// Get the underlying value and encode it
		optionalVal := fieldValue.MethodByName("MustGet").Call([]reflect.Value{})[0]
		optionalType := optionalVal.Type()

		return marshalField(finalBytes, optionalType, optionalVal, structField)
	}

	// If field is still a pointer, get rid of that now that we're past the optional checking
	fieldValue = reflect.Indirect(fieldValue)

	switch fieldValue.Kind() {
	case reflect.Uint8:
		newInt := uint8(fieldValue.Uint())
		finalBytes = append(finalBytes, newInt)
	case reflect.Uint16:
		newInt := uint16(fieldValue.Uint())
		finalBytes = append(finalBytes, util.Uint16ToBytes(newInt)...)
	case reflect.Uint32:
		newInt := uint32(fieldValue.Uint())
		finalBytes = append(finalBytes, util.Uint32ToBytes(newInt)...)
	case reflect.Uint64:
		finalBytes = append(finalBytes, util.Uint64ToBytes(fieldValue.Uint())...)
	case reflect.Array:
		switch fieldType.Elem().Kind() {
		case reflect.Uint8:
			// special case as byte-string
			for i := 0; i < fieldValue.Len(); i++ {
				finalBytes, err = marshalField(finalBytes, fieldType.Elem(), fieldValue.Index(i), structField)
				if err != nil {
					return finalBytes, err
				}
			}
		default:
			return finalBytes, fmt.Errorf("unimplemented array type %s", fieldType.Elem().Kind())
		}
	case reflect.Struct:
		finalBytes, err = marshalStruct(finalBytes, fieldType, fieldValue)
		if err != nil {
			return finalBytes, err
		}
	case reflect.Slice:
		finalBytes, err = marshalSlice(finalBytes, fieldType, fieldValue)
		if err != nil {
			return finalBytes, err
		}
	case reflect.String:
		// Strings get converted to []byte with a 4 byte size prefix
		strBytes := []byte(fieldValue.String())
		numBytes := uint32(len(strBytes))
		finalBytes = append(finalBytes, util.Uint32ToBytes(numBytes)...)

		finalBytes = append(finalBytes, strBytes...)
	case reflect.Bool:
		if fieldValue.Bool() {
			finalBytes = append(finalBytes, boolTrue)
		} else {
			finalBytes = append(finalBytes, boolFalse)
		}
	default:
		return finalBytes, fmt.Errorf("unimplemented type %s", fieldValue.Kind())
	}

	return finalBytes, nil
}

func marshalSlice(finalBytes []byte, t reflect.Type, v reflect.Value) ([]byte, error) {
	var err error

	// Slice/List is 4 byte prefix (number of items) and then serialization of each item
	// Get 4 byte length prefix
	numItems := uint32(v.Len())
	finalBytes = append(finalBytes, util.Uint32ToBytes(numItems)...)

	sliceKind := t.Elem().Kind()
	switch sliceKind {
	case reflect.Uint8: // same as byte
		// This is the easy case - already a slice of bytes
		finalBytes = append(finalBytes, v.Bytes()...)
	case reflect.Struct:
		for j := 0; j < int(numItems); j++ {
			currentStruct := v.Index(j)

			finalBytes, err = marshalStruct(finalBytes, currentStruct.Type(), currentStruct)
			if err != nil {
				return finalBytes, err
			}
		}
	}

	return finalBytes, nil
}
