package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	var result []int
	list := []int{3, 5, 1, 2, 4}
	result = algorithms.QuickSort(list)
	resultStr := fmt.Sprint(result)
	println(resultStr)
}
