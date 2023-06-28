package streamable

import (
	"reflect"
)

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// Error outputs the error message and satisfies the Error interface
func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "streamable: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "streamable: Unmarshal(nil " + e.Type.String() + ")"
}
