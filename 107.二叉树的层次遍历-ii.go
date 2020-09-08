/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
 func levelOrderHelper(nodes []*TreeNode) [][]int {
	if len(nodes) == 0 {
		return nil
	}

	newNodes := make([]*TreeNode, 0)
	vals := make([]int, len(nodes))
	for i, node := range nodes {
		vals[i] = node.Val

		if node.Left != nil {
			newNodes = append(newNodes, node.Left)
		}

		if node.Right != nil {
			newNodes = append(newNodes, node.Right)
		}
	}

	return append(levelOrderHelper(newNodes), vals)
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	return levelOrderHelper([]*TreeNode{root})
}
// @lc code=end

