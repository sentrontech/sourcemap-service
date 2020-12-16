package intmath

// Max returns the highest of two integers
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
