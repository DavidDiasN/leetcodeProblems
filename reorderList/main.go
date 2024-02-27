package main

func main() {
	head := &ListNode{5, nil}
	next := &ListNode{3, nil}
	head.Next = next
	reorderList(head)
}

func reorderList(head *ListNode) {
	nums := []int{}
	i := 0
	iter := head
	for iter != nil {
		nums = append(nums, iter.Val)
		iter = iter.Next
		i++
	}
	head = head.Next
	b := 1
	for b < i {
		head.Val = nums[i-1]
		if head.Next == nil {
			break
		}
		head.Next.Val = nums[b]
		head = head.Next.Next
		b++
		i--

	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}
