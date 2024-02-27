package main

func main() {

}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow2 := 0
	for slow2 != slow {
		slow = nums[slow]
		slow2 = nums[slow2]
	}
	return slow2
}
