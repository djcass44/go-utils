package mathutils

// Min returns the smallest of the 2 given integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the largest of the 2 given integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
