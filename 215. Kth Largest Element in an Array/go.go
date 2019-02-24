package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{534, 1234, 1, 124, 11, 4, 1, 612, 2}, 4))
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *intHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *intHeap) Pop() (i interface{}) {
	n := len(*h)
	i = (*h)[n-1]
	*h = (*h)[:n-1]
	return
}

func findKthLargest(nums []int, k int) int {
	hp := &intHeap{}
	for _, num := range nums {
		*hp = append(*hp, num)
	}
	heap.Init(hp)
	res := 0
	for i := 0; i < k; i++ {
		res = heap.Pop(hp).(int)
	}

	return res
}
