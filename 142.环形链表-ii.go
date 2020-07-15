/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
   slow,fast := head,head
   for fast!=nil && fast.Next != nil {
	  fast = fast.Next.Next
	  slow = slow.Next
	  if fast == slow {
		  break
	  }
   }

   if fast == nil || fast.Next == nil{
	   return nil
   }

   slow = head
   for slow != fast {
	   slow = slow.Next
	   fast = fast.Next
   }

   return fast
}
// @lc code=end

