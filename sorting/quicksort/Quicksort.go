package quicksort

import "fmt"

func Sort(toSort []int) {
	fmt.Println(findPivotIndex(toSort, 0, len(toSort)-1))
}

func quicksort(toSort []int, low int, high int) {
	if low > high {
		return
	}

	pivot := findPivotIndex(toSort, low, high)
	quicksort(toSort, 0, pivot)
	quicksort(toSort, pivot+1, high)
}

// Find the pivot point
func findPivotIndex(toSort []int, low int, high int) int {
	tmpPivot := toSort[high]
	tmpLower := low

	for i := low; i < high; i++ {
		if toSort[i] < tmpPivot {
			toSort[i], toSort[tmpLower] = toSort[tmpLower], toSort[i]
			tmpLower++
		}
	}
	toSort[tmpLower], toSort[high] = toSort[high], toSort[tmpLower]
	return tmpLower
}
