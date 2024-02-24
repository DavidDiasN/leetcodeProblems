package main

import "fmt"

func main() {
	ints1 := []int{1, 2, 3, 4}
	ints2 := []int{-1, 1, 0, -3, 3}
	ints3 := []int{55, 1, 3, 25, 1, 3, 25, 1, 3, 2, 1, 3, 2}

	fmt.Println(productExceptSelf(ints1))
	fmt.Println(productExceptSelf(ints2))
	fmt.Println(productExceptSelf(ints3))
}

func productExceptSelf(nums []int) []int {
	outputArray := make([]int, len(nums))

	total := 1
	for i := range nums {
		if nums[i] == 0 {
			total *= 1
		} else {
			total *= nums[i]
		}
	}

	for i := range outputArray {
		outputArray[i] = total
	}

	for i := range outputArray {
		if nums[i] == 0 {
			for j := range nums {
				if i != j {
					outputArray[j] = 0
				}
			}
		} else {
			outputArray[i] /= nums[i]
		}

	}
	return outputArray
}
