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
	l := len(nums)
	outputArray := make([]int, l)
	outputArray[0] = 1
	for i := 1; i < l; i++ {
		outputArray[i] = outputArray[i-1] * nums[i-1]
	}
	accu := 1
	for i := l - 1; i >= 0; i-- {
		outputArray[i] = outputArray[i] * accu
		accu = nums[i] * accu
	}
	return outputArray
}
