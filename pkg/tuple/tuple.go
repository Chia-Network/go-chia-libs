package tuple

import (
	"bytes"
	"encoding/json"
	"reflect"
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
// This works by creating a temp []interface{}, where each item in the slice is a pointer to the fields of the underlying
// tuple struct. Since it's pre-filled with pointers, unmarshal fills the pointers and checks the types
func (t *Tuple[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		t.isPresent = false
		return nil
	}

	v := reflect.ValueOf(&t.value).Elem()
	tmpMap := make([]interface{}, v.NumField())
	tmv := reflect.ValueOf(&tmpMap).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tmv.Index(i).Set(field.Addr())
	}
	err := json.Unmarshal(b, &tmpMap)
	if err != nil {
		return err
	}

	t.isPresent = true
	return nil
}
