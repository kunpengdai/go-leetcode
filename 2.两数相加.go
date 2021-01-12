/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	v := l1.Val + l2.Val
	res := &ListNode{
		Val: v % 10,
	}
	res.Next = addTwoNumbers(l1.Next, l2.Next)
	if v >= 10 {
		res.Next = addTwoNumbers(&ListNode{Val: v / 10}, res.Next)
	}

	return res
}

// @lc code=end

