/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	dfs(root, nil)
	return root
}

func dfs(node, next *Node) {
	if node != nil {
		node.Next = next
		dfs(node.Left, node.Right)
		if node.Next != nil {
			dfs(node.Right, node.Next.Left)
		}else{
			dfs(node.Right,nil)
		}
	}
}

// @lc code=end

