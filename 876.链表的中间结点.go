/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	fast,slow := head,head
	for fast!=nil && fast.Next!=nil{
		slow = slow.Next
		if fast.Next.Next!=nil{
			fast = fast.Next.Next
		}else{
			fast = nil
		}
	}
	return slow
}
// @lc code=end

