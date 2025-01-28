package util

import (
	"strconv"
)

// IsNumericInt returns true if the given string represents an integer
func IsNumericInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
