/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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

func reverseSlice(sl []int) {
	for i := 0; i < len(sl)/2; i++ {
		sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
	}
}

func reverse(res [][]int) [][]int {
	for i := range res {
		if i&1 == 1 {
			reverseSlice(res[i])
		}
	}

	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	return reverse(levelOrder(root))
}
// @lc code=end

