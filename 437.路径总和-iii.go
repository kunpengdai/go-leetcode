/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
 func pathSum(root *TreeNode, sum int) int {
	var res int
	pathSumHelper(root, &res, sum, []int{})
	return res
}

func pathSumHelper(root *TreeNode, res *int, sum int, nums []int) {

	if root == nil {
		return
	}

	nums = append(nums, root.Val)
	var tmp int
	for i := len(nums) - 1; i >= 0; i-- {
		tmp += nums[i]
		if tmp == sum {
			*res = *res + 1
		}
	}

	pathSumHelper(root.Left, res, sum, nums)
	pathSumHelper(root.Right, res, sum, nums)

	nums = nums[:len(nums)-1]
}
// @lc code=end

