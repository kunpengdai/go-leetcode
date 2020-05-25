/*
 * @lc app=leetcode.cn id=1145 lang=golang
 *
 * [1145] 二叉树着色游戏
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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

	node := dfs(root,x)
	leftNum := num(node.Left)
	rightNum := num(node.Right)
	if leftNum> n/2 || rightNum> n/2{
		return true
	}

	if (leftNum+rightNum)<n/2{
		return true
	}

	return false
}

func dfs(node *TreeNode,x int)*TreeNode{
	if node==nil{
		return nil
	}
	if node.Val == x{
		return node
	}

	ln := dfs(node.Left,x)
	if ln!=nil{
		return ln
	}

	return dfs(node.Right,x)
}



func num(node *TreeNode)int{
	if node==nil{
		return 0
	}
	return num(node.Left)+num(node.Right)+1
} 
// @lc code=end

