/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	flattenHelper(root)
}

func flattenHelper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := root.Right
	if root.Left == nil {
		flattenHelper(root.Right)
	} else {
		root.Right = root.Left
		flattenHelper(root.Left).Right = right
		flattenHelper(root.Right)
		root.Left = nil
	}

	node := root
	for node.Right != nil {
		node = node.Right
	}

	return node
}

// @lc code=end

