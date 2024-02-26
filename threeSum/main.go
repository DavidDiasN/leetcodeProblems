package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("yee")

}

func neetThreeSum(nums []int) [][]int {
	var result [][]int
	sort.Sort(sort.IntSlice(nums))

	for i, val := range nums {
		if i > 0 && val == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			total := val + nums[l] + nums[r]
			if total > 0 {
				r--
			} else if total < 0 {
				l++
			} else {
				result = append(result, []int{val, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

	}
	return result
}
