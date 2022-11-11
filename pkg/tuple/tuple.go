package tuple

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/samber/mo"
)

// Some builds a Tupe when value is present.
func Some[T any](value T) Tuple[T] {
	return Tuple[T]{
		isPresent: true,
		value:     value,
	}
}

// Tuple Wrapper for structs that encodes/decodes json as a tuple, rather than dict
type Tuple[T any] struct {
	isPresent bool
	value     T
}

// MarshalJSON custom marshaller for tuple wrapped structs
func (t Tuple[T]) MarshalJSON() ([]byte, error) {
	v := reflect.ValueOf(t.value)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return json.Marshal(values)
}

// UnmarshalJSON decodes Option from json.
func (t *Tuple[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		t.isPresent = false
		return nil
	}

	var tmpMap []interface{}
	err := json.Unmarshal(b, &tmpMap)
	if err != nil {
		return err
	}

	vp := reflect.ValueOf(&t.value)
	v := vp.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		setType(&field, tmpMap[i])
	}

	t.isPresent = true
	return nil
}

func setType(field *reflect.Value, fillValue interface{}) {
	if fillValue == nil {
		return
	}
	switch field.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
		// numeric looking things seem to unmarshal generically into float64
		if floated, ok := fillValue.(float64); ok {
			field.SetUint(uint64(floated))
		}
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.String:
		if stringed, ok := fillValue.(string); ok {
			field.SetString(stringed)
		}
	case reflect.Struct:
		// The only nested struct we support for now is Option[T]
		actualType := field.Type().Name()
		switch actualType {
		case "Option[string]":
			os := mo.Some(fillValue.(string))
			osrp := reflect.ValueOf(&os)
			osr := osrp.Elem()
			field.Set(osr)
		}
		// @TODO handle other types
	}
}
