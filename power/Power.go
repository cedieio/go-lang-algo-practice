package power

// Power accepts num and the power, returns the num and the power
func Power(num int, pow int) int {
	if pow == 0 {
		return 1
	}

	// 5 ^ 3 call stack would look like
	// 5 * Power(5, 3 - 1) = 5 * 25 since pow(5, 3-1)  would return 25
	// 5 * Power(5, 2 - 1) = 5 * 5 since pow(5, 2-1) would return 5
	// 5 * Power(5, 1 - 1) = 5 * 1 since pow(5, 1-1) would return 1
	return num * Power(num, pow-1)
}

// 8 * func(8 - 2)
// 6 * func(6 - 2) 48
// 4 * func(4 - 2) 8
// 2 * func(2 - 2) 2
