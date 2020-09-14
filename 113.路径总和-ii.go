/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func pathSum(root *TreeNode, sum int) [][]int {

	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	pathSumHelper(root, sum, 0, []int{}, &res)
	return res
}

func pathSumHelper(root *TreeNode, sum int, sumNow int, sl []int, res *[][]int) {

	if root.Left == nil && root.Right == nil {
		if sumNow+root.Val == sum {
			sl = append(sl, root.Val)
			dst := make([]int, len(sl))
			copy(dst, sl)
			*res = append(*res, dst)
			sl = sl[:len(sl)-1]
		}
	}

	if root.Left != nil {
		sl = append(sl, root.Val)
		pathSumHelper(root.Left, sum, sumNow+root.Val, sl, res)
		sl = sl[:len(sl)-1]
	}

	if root.Right != nil {
		sl = append(sl, root.Val)
		pathSumHelper(root.Right, sum, sumNow+root.Val, sl, res)
		sl = sl[:len(sl)-1]
	}
}

// @lc code=end

