package mergesort

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := arr[:mid]
	rightArr := arr[mid:]
	leftArr = Sort(leftArr)
	rightArr = Sort(rightArr)
	return merge(leftArr, rightArr)
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		merged[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		merged[k] = right[j]
		k++
		j++
	}
	return merged

}
