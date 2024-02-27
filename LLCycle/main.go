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
