/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	p1,p2 := head,head
	for p1!=nil && p2!=nil{

		for p1!=nil{
			if p1.Val >= x{
				break
			}
			p1 = p1.Next
		}

		for p2=p1.Next;p2!=nil;p2=p2.Next{
			if p2.Val < x{
				break
			}
		}

		if p1!=nil && p2!=nil{
			p1.Val,p2.Val = p2.Val,p1.Val
			p1 = p1.Next
			p2 = p2.Next
		}
	}

	return head
}
// @lc code=end

