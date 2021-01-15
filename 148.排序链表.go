/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sortListHelper(head, nil)
}

func sortListHelper(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	mid := findMedianNode(head, tail)
	return merge(sortListHelper(head, mid), sortListHelper(mid, tail))
}

func findMedianNode(head, tail *ListNode) *ListNode {
	fast, slow := head, head
	if fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != tail {
			fast = fast.Next
		}
	}

	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp := dummyHead
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}

	return dummyHead.Next
}

// @lc code=end

