package main

func main() {

}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleBetter(head *ListNode) bool {
	if head == nil {
			return false
	}
	slow, fast := head, head.Next

	for fast != nil {
			if fast.Next == nil {
					return false
			}
			if slow == fast {
					return true
			}
			slow = slow.Next
			fast = fast.Next.Next
	}
	return false
}