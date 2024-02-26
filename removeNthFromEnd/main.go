package main

import "fmt"

func main() {
	myList := &ListNode{}
	myNums := [3]int{1, 2, 3}
	myPlace := myList
	for i, val := range myNums {

		myPlace.Val = val
		if i == len(myNums)-1 {
			break
		}
		nextNode := &ListNode{}
		myPlace.Next = nextNode
		myPlace = nextNode
	}
	printLL(myList)

	printLL(removeNthFromEnd(myList, 2))

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	b, a := head, head
	l := 0
	for i := n; i > 0; i-- {
		if a.Next != nil {
			l++
			a = a.Next
		} else {
			break
		}
	}

	for a != nil {
		if a.Next == nil {
			if b == a {
				return nil
			} else if head == b && b.Next == a {
				if n == 1 {
					head.Next = nil
					return head
				}
				return a

			} else if head == b && b.Next != a {
				if n == l+1 {
					head = head.Next
					return head
				}
				b.Next = b.Next.Next
				return head
			} else {
				b.Next = b.Next.Next
				return head
			}

		} else {
			l++
			b = b.Next
			a = a.Next
		}

	}

	return nil

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLL(myList *ListNode) {
	for myList != nil {
		fmt.Println(myList.Val)
		myList = myList.Next
	}
}
