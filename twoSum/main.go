package main

import (
	"fmt"
)

func main() {
	//	result := myTwoSum([]int{1, 2, 3, 4, 5}, 1)
	//	fmt.Printf("Result: %d\n", result)
	result := neetTwoSum([]int{1, 2, 3, 4, 5}, 9)
	fmt.Printf("Result: %d\n", result)
}

func myTwoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return []int{0, 0}
	}
	j := 1
	for i := 0; i <= numsLen-2; i++ {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		} else {
			if j < numsLen-1 {
				j++
				i--
			} else {
				j = i + 2
			}
		}
	}

	return []int{0, 0}
}

func neetTwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, ival := range nums {
		diff := target - ival
		index, has := m[diff]
		if has {
			return []int{index, i}
		} else {
			m[nums[i]] = i
		}

	}
	return []int{0, 0}
}
