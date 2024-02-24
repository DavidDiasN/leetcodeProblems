package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1.Next == nil && l2.Next == nil {
		total := l1.Val + l2.Val
		if total < 10 {
			return &ListNode{Val: total, Next: nil}
		} else {
			l1.Val = total / 10
			l2.Val = total % 10
			l2.Next = l1
			return l2
		}

	}

	myListNodeOne := &ListNode{Val: 0, Next: nil}
	carryOver := 0
	iterateMyList := myListNodeOne

	for l1 != nil || l2 != nil {
		lastCarry := carryOver
		carryOver = 0
		var val1 int
		var val2 int

		if l1 == nil {
			val1 = 0
			val2 = l2.Val
		} else if l2 == nil {
			val2 = 0
			val1 = l1.Val
		} else {
			val1 = l1.Val
			val2 = l2.Val
		}

		total := val1 + val2 + lastCarry

		if total > 9 {
			carryOver = 1
			iterateMyList.Val = total % 10
		} else {
			iterateMyList.Val = total
		}

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			nextNode := &ListNode{}
			iterateMyList.Next = nextNode
			iterateMyList = nextNode
		} else if carryOver != 0 {
			nextNode := &ListNode{Val: carryOver, Next: nil}
			iterateMyList.Next = nextNode
			iterateMyList = nextNode
		}
	}

	return myListNodeOne
}

// MULTIPLY BY PLACE, ITERATE THROUGH THEM SEPEATELY, THEN ITERATE THROUGH THE STRING YOU MAKE I GUESS
