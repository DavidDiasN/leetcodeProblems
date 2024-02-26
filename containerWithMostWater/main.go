package main

import (
	"fmt"
)

func main() {
	fmt.Println("lol")
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxA := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		maxA = max(maxA, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxA

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
