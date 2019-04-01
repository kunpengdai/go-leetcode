package main

import "fmt"

type item struct {
	x, min int
}

type MinStack struct {
	slice []item
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.slice) > 0 && this.GetMin() < min {
		min = this.GetMin()
	}
	this.slice = append(this.slice, item{
		x:   x,
		min: min,
	})
}

func (this *MinStack) Pop() {
	this.slice = this.slice[:len(this.slice)-1]
}

func (this *MinStack) Top() int {
	return this.slice[len(this.slice)-1].x
}

func (this *MinStack) GetMin() int {
	return this.slice[len(this.slice)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
