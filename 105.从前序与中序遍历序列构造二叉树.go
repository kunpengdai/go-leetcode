/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}

	return nil
}

// @lc code=end

