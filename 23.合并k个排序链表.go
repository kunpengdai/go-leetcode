/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type nodes []*ListNode

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}

func (n *nodes) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *nodes) Push(x interface{}) {
	*n = append(*n, x.(*ListNode))
}

func (n *nodes) Pop() interface{} {
	res := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {

	head := &ListNode{}
	pos := head

	ns := make(nodes, 0)
	heap.Init(&ns)
	for _, node := range lists {
		if node != nil {
			heap.Push(&ns, node)
		}
	}

	for len(ns) != 0 {
		minV := heap.Pop(&ns)
		minNode := minV.(*ListNode)
		pos.Next = minNode
		pos = pos.Next
		if minNode.Next != nil {
			heap.Push(&ns, minNode.Next)
		}
	}

	return head.Next
}

// @lc code=end

