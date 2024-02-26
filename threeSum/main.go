package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("yee")
	twi := []int{5, 2, 1}
	sort.Sort(sort.IntSlice(twi))
	fmt.Println(twi)

}

func threeSum(nums []int) [][]int {
	var result [][]int
	lastPoint := 0
	m := make(map[int]int)
	sort.Sort(sort.IntSlice(nums))

	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		}

	}

}
