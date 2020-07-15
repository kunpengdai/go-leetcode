/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}

	slow,fast := head,head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
// @lc code=end

