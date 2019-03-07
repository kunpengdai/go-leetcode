package main

import "fmt"

func main() {
	t0 := &TreeNode{Val: 0}
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 7}
	t8 := &TreeNode{Val: 8}

	t3.Left = t5
	t3.Right = t1

	t5.Left = t6
	t5.Right = t2

	t1.Left = t0
	t1.Right = t8

	t2.Left = t7
	t2.Right = t4

	fmt.Println(lowestCommonAncestor(t3, t5, t1).Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
