package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	print(mergeTwoLists(l1, l2))
}

func print(l1 *ListNode) {
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res, p, p1, p2 := &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		res = l1
		p1 = l1.Next
		p2 = l2
	} else {
		res = l2
		p1 = l1
		p2 = l2.Next
	}
	p = res
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return res

}
