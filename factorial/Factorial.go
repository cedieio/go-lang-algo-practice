package factorial

// Factorial accepts int and computes the factorial
func Factorial(num int) int {
	if num == 0 {
		return 1
	}

	// 5 would look like
	// 5 * Factorial(4) 120
	// 4 * Factorial(3) 24
	// 3 * Factorial(2) 6
	// 2 * Factorial(1) 2
	// 1 * Factorial(0) 1
	// It's a stack so it would compute 1 * Factorial(0) which is 1 * 1 and so on
	return num * Factorial(num-1)
}
