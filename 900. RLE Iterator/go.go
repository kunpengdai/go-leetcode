package main

import "fmt"

func main() {
	obj := Constructor([]int{1, 2, 3, 3})
	fmt.Println(obj.Next(3))
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(1))

}

type RLEIterator struct {
	index  int
	ithNum int
	nums   []int
}

func Constructor(A []int) RLEIterator {
	ret := RLEIterator{
		nums: A,
	}
	return ret
}

func (this *RLEIterator) Next(n int) int {
	for n > 0 && this.ithNum < len(this.nums) {
		if this.nums[this.ithNum]-this.index >= n {
			this.index += n
			return this.nums[this.ithNum+1]
		}
		n -= (this.nums[this.ithNum] - this.index)
		this.index = 0
		this.ithNum += 2
	}
	return -1
}
