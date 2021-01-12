/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

type sumNode struct {
	Val     int
	Left    *sumNode
	Right   *sumNode
	maxSum  int
	maxSide int
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	node := copyNode(root)
	maxPathSumHepler(node)
	res := root.Val
	dfs(node, &res)
	return res
}

func copyNode(root *TreeNode) *sumNode {
	if root == nil {
		return nil
	}

	return &sumNode{
		Val:   root.Val,
		Left:  copyNode(root.Left),
		Right: copyNode(root.Right),
	}
}

func dfs(root *sumNode, res *int) {
	if root == nil {
		return
	}

	*res = max(root.maxSum, *res)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxPathSumHepler(root *sumNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		root.maxSide = root.Val
		root.maxSum = root.Val
		return
	}

	var left, right int
	if root.Left != nil {
		maxPathSumHepler(root.Left)
		left = root.Left.maxSide
	}

	if root.Right != nil {
		maxPathSumHepler(root.Right)
		right = root.Right.maxSide
	}

	root.maxSide = root.Val + max(0, max(left, right))
	root.maxSum = root.Val + max(0, left) + max(0, right)
}

// @lc code=end

