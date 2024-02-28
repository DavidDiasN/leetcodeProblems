package main

func man() {

}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	m := make(map[*Node]*Node)
	iter := &Node{}
	saveIter := iter
	headCopy := head
	for headCopy != nil {
		iter.Val = headCopy.Val
		m[headCopy] = iter
		if head.Next != nil {
			nextNode := &Node{}
			iter.Next = nextNode
			iter = iter.Next
			headCopy = headCopy.Next
		}
	}
	lastIter := saveIter
	for head != nil {
		lastIter.Random = m[head.Random]
		lastIter = lastIter.Next
		head = head.Next
	}
	return saveIter
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
