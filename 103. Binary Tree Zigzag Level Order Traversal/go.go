package main

import "fmt"

func main() {
	t0 := &TreeNode{Val: 0}
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 7}
	t8 := &TreeNode{Val: 8}

	t3.Left = t5
	t3.Right = t1

	t5.Left = t6
	t5.Right = t2

	t1.Left = t0
	t1.Right = t8

	t2.Left = t7
	t2.Right = t4

	fmt.Println(zigzagLevelOrder(t3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	bfs := make([][]*TreeNode, 0)
	if root == nil {
		return [][]int{}
	}
	q1, q2 := make([]*TreeNode, 1), make([]*TreeNode, 0)
	q1[0] = root

	for len(q1) > 0 {
		bfs = append(bfs, q1)
		for _, item := range q1 {
			if item.Left != nil {
				q2 = append(q2, item.Left)
			}
			if item.Right != nil {
				q2 = append(q2, item.Right)
			}
		}
		q1 = q2
		q2 = make([]*TreeNode, 0)
	}

	fmt.Println("bfs:", bfs)

	res := make([][]int, 0)
	for i := range bfs {
		is := make([]int, 0)
		for j := range bfs[i] {
			if i&1 == 0 {
				is = append(is, bfs[i][j].Val)
			} else {
				is = append(is, bfs[i][len(bfs[i])-1-j].Val)
			}
		}
		res = append(res, is)
	}

	return res
}
