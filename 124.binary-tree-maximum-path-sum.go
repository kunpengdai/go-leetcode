/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil{
		return 0
	}
	res := root.Val
	helper(root,&res)
	return res
}

func helper(node *TreeNode,res *int)int{
	if node == nil{
		return 0
	}
	l := max(helper(node.Left,res),0)
	r := max(helper(node.Right,res),0)
	*res = max(*res,l+r+node.Val)
	return max(l+node.Val,r+node.Val)
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

