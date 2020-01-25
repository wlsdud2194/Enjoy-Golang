package main

import "fmt"

func findMinValue(arr []int, startIdx int) (int, int) {
	var (
		idx      int = startIdx
		minValue int = arr[startIdx]
	)

	for i := startIdx + 1; i < len(arr); i++ {
		if minValue >= arr[i] {
			idx = i
			minValue = arr[i]
		}
	}

	return idx, minValue
}

func selection_sort(arr []int) []int {
	var (
		temp int
	)

	for i := 0; i < len(arr)-1; i++ {
		j, _ := findMinValue(arr, i)

		if arr[i] > arr[j] {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	return arr
}

func ExampleSort() {
	var (
		arr1 = []int{9, 5, 3, 8, 2, 7}
	)

	arr1 = selection_sort(arr1)

	fmt.Println(arr1)
	// Output:
	// [2 3 5 7 8 9]
}
