/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	values := make([]int, 0)
	for head!=nil{
		values = append(values,head.Val)
		head = head.Next
	}

	for i:=0;i<len(values)/2;i++{
		if values[i]!=values[len(values)-i-1]{
			return false
		}
	}

	return true
}
// @lc code=end
