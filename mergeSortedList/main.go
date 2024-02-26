package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	i1, i2 := list1, list2
	newLL := &ListNode{}
	newLLHead := newLL
	for i1 != nil || i2 != nil {
		if i1 == nil {
			fmt.Println("i1 == nil")
			newLL.Val = i2.Val
			newLL.Next = i2.Next
			return newLLHead
		} else if i2 == nil {
			newLL.Val = i1.Val
			newLL.Next = i1.Next
			fmt.Println("i2 == nil")
			return newLLHead
		}
		if i1.Val > i2.Val {
			newLL.Val = i2.Val
			i2 = i2.Next
		} else {
			newLL.Val = i1.Val

			i1 = i1.Next
		}
		newLL.Next = &ListNode{}
		newLL = newLL.Next

	}
	fmt.Println("We got here")
	return newLLHead
}
