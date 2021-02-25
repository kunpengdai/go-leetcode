/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
    ans, pre := math.MaxInt64, -1
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if pre != -1 && node.Val-pre < ans {
            ans = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return ans
}
// @lc code=end

