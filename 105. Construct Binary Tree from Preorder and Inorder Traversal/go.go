package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))

	fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := new(TreeNode)
	if len(preorder) == 0 {
		return nil
	}
	fisrtNum := preorder[0]
	index := 0
	for i, item := range inorder {
		if item == fisrtNum {
			index = i
			break
		}
	}
	root.Val = fisrtNum

	if index == 0 {
		root.Left = nil
		root.Right = buildTree(preorder[1:], inorder[1:])
	} else if index == len(preorder)-1 {
		root.Left = buildTree(preorder[1:index+1], inorder[0:index])
		root.Right = nil
	} else {
		root.Left = buildTree(preorder[1:index+1], inorder[0:index])
		root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	}

	return root
}
