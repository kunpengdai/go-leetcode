/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return pathSumHelper(root, sum, 0)
}

func pathSumHelper(root *TreeNode, sum int, sumNow int) bool {
	if root.Left == nil && root.Right == nil {
		return sumNow+root.Val == sum
	}

	return root.Left != nil && pathSumHelper(root.Left, sum, sumNow+root.Val) ||
		root.Right != nil && pathSumHelper(root.Right, sum, sumNow+root.Val)
}

// @lc code=end

