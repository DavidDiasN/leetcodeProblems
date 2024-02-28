package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(-2)
	fmt.Println("checking state")
	fmt.Println(obj.Val)
	obj.Push(0)
	fmt.Println("checking state")
	fmt.Println(obj.Val)
	obj.Push(-3)
	fmt.Println("checking state")
	fmt.Println(obj.Val)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())

}

type MinStack struct {
	Val  int
	Prev *MinStack
	Min  int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.Prev == nil {
		this.Min = val
		this.Val = val
		this.Prev = &MinStack{}
	} else {
		saveThis := &MinStack{Val: this.Val, Prev: this.Prev, Min: this.Min}
		this.Val = val
		this.Prev = saveThis
		this.Min = val
		if saveThis.Min < this.Val {
			this.Min = saveThis.Min
		}
	}
}

func (this *MinStack) Pop() {
	this.Val = this.Prev.Val
	this.Min = this.Prev.Min
	this.Prev = this.Prev.Prev
}

func (this *MinStack) Top() int {
	return this.Val
}

func (this *MinStack) GetMin() int {
	return this.Min
}
