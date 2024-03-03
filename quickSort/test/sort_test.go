package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	arr1 := generateArray()
	arr2 := generateArray()
	quickSort(arr1, 0, len(arr1)-1)
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("BenchmarkSort")
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func BenchmarkConSort(b *testing.B) {
	arr1 := generateArray()
	arr2 := generateArray()

	go quickSort(arr1, 0, len(arr1)/4)
	go quickSort(arr1, (len(arr1)/4)+1, (len(arr1) / 2))

	go quickSort(arr1, (len(arr1)/2)+1, (len(arr1)/4)*3)
	go quickSort(arr1, ((len(arr1)/4)*3)+1, len(arr1)-1)

	go quickSort(arr2, 0, len(arr2)/4)
	go quickSort(arr2, (len(arr2)/4)+1, (len(arr2) / 2))

	go quickSort(arr2, (len(arr2)/2)+1, (len(arr2)/4)*3)
	go quickSort(arr2, ((len(arr2)/4)*3)+1, len(arr2)-1)
	fmt.Printf("BenchmarkConSort")
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func generateArray() []int {
	size := 1000
	res := []int{}
	for i := 0; i < size; i++ {
		res = append(res, rand.Intn(500))
	}
	return res
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
