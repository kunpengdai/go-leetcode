/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type item struct {
	val int
	min int
}

type MinStack struct {
	items []item
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items: make([]item, 0, 100),
	}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.items) != 0 && this.items[len(this.items)-1].min < min {
		min = this.items[len(this.items)-1].min
	}

	this.items = append(this.items, item{
		val: x,
		min: min,
	})
}

func (this *MinStack) Pop() {
	if len(this.items) != 0 {
		this.items = this.items[:len(this.items)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.items) != 0 {
		return this.items[len(this.items)-1].val
	}

	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.items) != 0 {
		return this.items[len(this.items)-1].min
	}

	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

