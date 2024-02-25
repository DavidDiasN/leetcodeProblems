package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
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
