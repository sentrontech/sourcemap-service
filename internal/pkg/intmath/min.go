package intmath

// Min returns the lowest of two integers
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
