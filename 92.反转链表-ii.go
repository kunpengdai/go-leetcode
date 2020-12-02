/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	cnt := n - m + 1
	stack := make([]*ListNode, cnt)

	start := &ListNode{
		Next: head,
	}

	m--
	p := start
	for m != 0 {
		p = p.Next
		m--
	}

	q := p
	for i := 0; i < cnt; i++ {
		stack[i] = p.Next
		p = p.Next
	}
	t := p.Next

	for i := len(stack) - 1; i >= 0; i-- {
		q.Next = stack[i]
		q = q.Next
	}

	q.Next = t
	return start.Next
}

// @lc code=end

