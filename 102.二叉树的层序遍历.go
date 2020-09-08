/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	list := []*TreeNode{root}
	for len(list) > 0 {

		vals := make([]int, 0)
		newList := make([]*TreeNode, 0)

		for _, item := range list {
			vals = append(vals, item.Val)
			if item.Left != nil {
				newList = append(newList, item.Left)
			}

			if item.Right != nil {
				newList = append(newList, item.Right)
			}
		}

		list = newList
		res = append(res, vals)
	}

	return res
}

// @lc code=end

