package bubblesort

func Sort(toSort []int) []int {
	arrCopy := toSort

	for i := len(toSort) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arrCopy[j] > arrCopy[j+1] {
				arrCopy[j], arrCopy[j+1] = arrCopy[j+1], arrCopy[j]
			}
		}
	}
	return arrCopy
}
