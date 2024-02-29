package main

func main() {

}

type LRUCache struct {
	Cache    map[int]*Node
	Capacity int
	Head     *Node
	Tail     *Node
}

type Node struct {
	CacheVal int
	CacheKey int
	Next     *Node
	Prev     *Node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev, node.Next = nil, nil
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.Head.Prev, this.Head
	prev.Next, node.Prev = node, prev
	node.Next, next.Prev = next, node
}

func Constructor(capacity int) LRUCache {
	tail := &Node{}
	head := &Node{}
	tail.Next, head.Prev = head, tail

	return LRUCache{
		Cache:    map[int]*Node{},
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.insert(node)
	return node.CacheVal
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		this.remove(node)
	}

	this.Cache[key] = &Node{CacheVal: value, CacheKey: key}
	this.insert(this.Cache[key])

	if len(this.Cache) > this.Capacity {
		lru := this.Tail.Next
		this.remove(lru)
		delete(this.Cache, lru.CacheKey)
	}

}
