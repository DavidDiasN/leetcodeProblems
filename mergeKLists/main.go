package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	for i := len(lists) - 1; i >= 0; i-- {
		if lists[i] == nil {
			lists = remove(lists, i)
		}
	}

	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	for i := len(lists) - 1; i > 0; i-- {
		merged := &ListNode{}
		saveMergedHead := merged
		sorted := lists[i]
		list1 := lists[i-1]
		for sorted != nil || list1 != nil {
			if sorted == nil {
				merged.Next = list1
				break
			} else if list1 == nil {
				merged.Next = sorted
				break
			}

			if sorted.Val < list1.Val {
				merged.Val = sorted.Val
				sorted = sorted.Next
			} else {
				merged.Val = list1.Val
				list1 = list1.Next
			}

			if sorted == nil || list1 == nil {
				continue
			}
			nextNode := &ListNode{}
			merged.Next = nextNode
			merged = merged.Next

		}
		lists[i-1] = saveMergedHead
	}

	return lists[0]
}

func remove(slice []*ListNode, s int) []*ListNode {
	return append(slice[:s], slice[s+1:]...)
}
