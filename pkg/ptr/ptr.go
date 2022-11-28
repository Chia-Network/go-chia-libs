package ptr

// Pointer convert any input value to pointer
func Pointer[K any](val K) *K {
	return &val
}
