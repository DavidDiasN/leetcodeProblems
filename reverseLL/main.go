package main

func main() {
    
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head

	savep2Next := p2.Next

	for p2 != nil {
		savep2Next = p2.Next
		if p2 == head {
			p2.Next = nil
		} else {
			p2.Next = p1
		}
		p1 = p2
		p2 = savep2Next
	}
	return p1

}
