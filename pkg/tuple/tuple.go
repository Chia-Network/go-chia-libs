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
		switch v.Field(i).Kind() {
		case reflect.String:
			if stringed, ok := tmpMap[i].(string); ok {
				v.Field(i).SetString(stringed)
			}
			// @TODO handle other types
		}
	}

	t.isPresent = true
	return nil
}
