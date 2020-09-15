/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
 func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var res int
	sumNumberHelper(root, &res, 0)
	return res
}

func sumNumberHelper(node *TreeNode, res *int, sumNow int) {
	if node.Left == nil && node.Right == nil {
		*res += sumNow*10 + node.Val
	}

	if node.Left != nil {
		sumNumberHelper(node.Left, res, sumNow*10+node.Val)
	}

	if node.Right != nil {
		sumNumberHelper(node.Right, res, sumNow*10+node.Val)
	}
}
// @lc code=end