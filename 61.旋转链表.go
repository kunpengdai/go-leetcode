/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}
	tail.Next = head

	k = length - (k % length)
	if k == 0 {
		return head
	}

	for i := 0; i < k; i++ {
		head = head.Next
		tail = tail.Next
	}

	tail.Next = nil
	return head
}
// @lc code=end

