package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1}))

}

func containsDuplicate(nums []int) bool {
	s := make(map[int]bool)

	for _, val := range nums {
		_, ok := s[val]
		if ok {
			return true
		}
		s[val] = true
	}

	return false
}
