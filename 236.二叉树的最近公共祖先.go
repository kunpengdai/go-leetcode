/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return root
	}

	path2p := []*TreeNode{root}
	path2q := []*TreeNode{root}
	findPath(root, p, &path2p)
	findPath(root, q, &path2q)

	res := path2p[0]
	for i := 1; i < len(path2p) && i < len(path2q); i++ {
		if path2p[i] != path2q[i] {
			break
		}
		res = path2p[i]
	}

	return res
}

func findPath(root, p *TreeNode, path *[]*TreeNode) bool {

	if root == nil {
		return false
	}

	*path = append(*path, root)
	if root == p {
		return true
	}

	if findPath(root.Left, p, path) {
		return true
	}

	if findPath(root.Right, p, path) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}

// @lc code=end

