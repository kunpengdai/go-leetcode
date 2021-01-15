/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var cntA, cntB int
	tempA, tempB := headA, headB
	for tempA != nil {
		cntA++
		tempA = tempA.Next
	}

	for tempB != nil {
		cntB++
		tempB = tempB.Next
	}

	if cntA > cntB {
		for i := 0; i < cntA-cntB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < cntB-cntA; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

// @lc code=end

