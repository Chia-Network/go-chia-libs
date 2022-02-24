package ptr

// IntPtr returns a pointer for the provided int
func IntPtr(i int) *int {
	return &i
}
