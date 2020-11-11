/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head,nil)
}

func buildTree(l,r *ListNode)*TreeNode {

	if l==r {
		return nil
	}

	mid := findMedian(l,r)
	return &TreeNode{
		Val:mid.Val,
		Left:buildTree(l,mid),
		Right:buildTree(mid.Next,r),
	}
}

func findMedian(l,r *ListNode) *ListNode {

	slow, fast := l, l
	for fast != r && fast.Next != r {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// @lc code=end

