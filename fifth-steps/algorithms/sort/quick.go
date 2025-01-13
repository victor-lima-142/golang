package algorithms

func QuickSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength <= 1 {
		return arr
	}

	pivotPosition := arrLength / 2
	pivot := arr[pivotPosition]
	left := []int{}
	right := []int{}

	for i, v := range arr {
		if i == pivotPosition {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
