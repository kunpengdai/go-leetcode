/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if n==0 {
		return head
	}

	tmp := &ListNode{
		Next:head,
	}
	p1,p2 := tmp,tmp
	for i := 0;i<=n;i++{
		p2 = p2.Next
	}

	for p2!=nil{
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next
	return tmp.Next
}
// @lc code=end

