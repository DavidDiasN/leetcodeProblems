package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := range numbers {
		diff := target - numbers[i]
		if _, ok := m[diff]; ok {
			return []int{m[diff] + 1, i + 1}
		} else {
			m[numbers[i]] = i
		}
	}
	return []int{}
}

func revisedTwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if check := numbers[l]+numbers[r] < target; check {
			l++
			continue
		} else if check := numbers[l]+numbers[r] > target; check {
			r--
			continue
		} else {
			return []int{l + 1, r + 1}
		}

	}
	return []int{-1, -1}

}
