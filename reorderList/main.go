package main

func main() {
	head := &ListNode{5, nil}
	next := &ListNode{3, nil}
	head.Next = next
	reorderList(head)
}

func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	// find positions
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse second half
	second := slow.Next
	var prev *ListNode
	prev = nil
	slow.Next = nil
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}
	// reorder the list
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}
