package euclid

// GenerateGCD uses a and b to get it's greatest common denominator, uses a for loop
func GenerateGCD(a int, b int) int {
	aCopy := a
	bCopy := b
	remainder := aCopy % bCopy
	for remainder > 0 {
		aCopy = bCopy
		bCopy = remainder
		remainder = aCopy % bCopy
	}

	return bCopy
}

// GenerateGCDRecursive uses inputs a and b to produce greatest common denominator, uses recursion
func GenerateGCDRecursive(a int, b int) int {
	if a%b == 0 {
		return b
	}

	return GenerateGCDRecursive(b, a%b)
}
