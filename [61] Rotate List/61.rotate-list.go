/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head==nil||head.Next==nil||k==0{
		return head
	}
	cnt,h := 1,head
	for h.Next!=nil{
		cnt++
		h = h.Next
	}
	k%=cnt
	if k==0{
		return head
	}
	h.Next = head
	k = cnt-k
	h = head
	for i:=0;i<k;i++{
		h = h.Next
	}
	p := h
	for p.Next != h{
		p = p.Next
	}
	p.Next = nil
	return h
}


