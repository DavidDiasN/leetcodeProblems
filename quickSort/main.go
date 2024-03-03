package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 5, 2, 10, 33, 4, 5, 2, 3, 4, 60, 34, 25, 66, 40, 55}
	arr2 := []int{4, 2, 6, 1, 6, 2, 7, 1, 8, 5, 9, 3, 8, 5, 86, 6, 2}
	fmt.Println(arr1)
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(arr2)
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Println(arr2)

}

func quickSort(arr []int, start, end int) {
	if end <= start {
		return
	}

	pivotIndex := partition(arr, start, end)
	quickSort(arr, start, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end)

}

func partition(arr []int, start, end int) int {
	ci, sm := start, start-1
	pv := arr[end]

	for ci <= end {
		if ci == end {
			sm++
			arr[ci], arr[sm] = arr[sm], arr[ci]
			return sm
		}
		if arr[ci] <= pv {
			sm++
			arr[ci], arr[sm] = arr[sm], arr[ci]
		}
		ci++
	}

	return sm
}
