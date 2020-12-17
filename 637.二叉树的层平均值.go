/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return nil
	}

	res := make([]float64, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var sum float64
		newNodes := make([]*TreeNode,0, len(nodes))
		for _, node := range nodes {
			sum += float64(node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}

			res = append(res, sum/float64(len(nodes)))
			nodes = newNodes
		}
	}

	return res
}

// @lc code=end

