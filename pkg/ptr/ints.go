package ptr

// IntPtr returns a pointer for the provided int
func IntPtr(i int) *int {
	return Pointer(i)
}

// Uint32Ptr returns a pointer for the provided uint32
func Uint32Ptr(i uint32) *uint32 {
	return Pointer(i)
}

// Uint64Ptr returns a pointer for the provided uint64
func Uint64Ptr(i uint64) *uint64 {
	return Pointer(i)
}
