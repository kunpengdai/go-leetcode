/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (33.64%)
 * Likes:    1023
 * Dislikes: 85
 * Total Accepted:    200.3K
 * Total Submissions: 587.5K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 1->1->1->2->3
 * Output: 2->3
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	h := &ListNode{Next: head,Val:head.Val-1}
	p:=h
	for p!=nil{
		if p.Next == nil || p.Next.Next == nil{
			break
		}

		if p.Next.Val == p.Next.Next.Val {
			t:= p.Next.Next
			for t != nil{
				if t.Val!=p.Next.Val{
					break
				}
				t=t.Next
			}
			p.Next = t
		}else{
			 p = p.Next
		}
	}

	return h.Next
}

