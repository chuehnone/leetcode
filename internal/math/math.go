package math

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
