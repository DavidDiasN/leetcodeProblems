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


 func neetMergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
			return nil
	}

	for len(lists) > 1 {
			mergedLists := []*ListNode{}
			for i := 0; i < len(lists); i+=2 {
					l1 := lists[i]
					var l2 *ListNode
					if i+1 <= len(lists)-1 {
							l2 = lists[i+1]
							mergedLists = append(mergedLists, mergeLists(l1, l2))
					} else {
							mergedLists = append(mergedLists, l1)
					}
			}
			lists = mergedLists
	}
	return lists[0]

}

func mergeLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
			return l2
	} else if l2 == nil {
			return l1
	}
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil || l2 != nil {
			if l1 == nil {
					dummy.Next = l2
					break
			} else if l2 == nil {
					dummy.Next = l1
					break
			}

			if l1.Val < l2.Val {
					dummy.Val = l1.Val
					l1 = l1.Next
			} else {
					dummy.Val = l2.Val
					l2 = l2.Next
			}

			if l1 == nil || l2 == nil {
					continue
			}
			nextDummy := &ListNode{}
			dummy.Next = nextDummy
			dummy = dummy.Next
	}

	return tail
}