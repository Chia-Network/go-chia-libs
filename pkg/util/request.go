package util

import (
	"github.com/google/uuid"
)

// GenerateRequestID generates a random string to use as a request ID
func GenerateRequestID() string {
	return uuid.New().String()
}
