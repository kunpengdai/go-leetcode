package main

import "fmt"

func main() {
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 7}

	print(sortList(l2))

}

func print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
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
func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func quickSort(head, tail *ListNode) {
	if head != tail {
		pt := partation(head, tail)
		quickSort(head, pt)
		quickSort(pt.Next, tail)
	}
}

func partation(head, tail *ListNode) *ListNode {
	if head == tail {
		return head
	}
	plow, phigh := head, head.Next
	for phigh != nil {
		if phigh.Val < head.Val {
			plow = plow.Next
			plow.Val, phigh.Val = phigh.Val, plow.Val
		}
		phigh = phigh.Next
	}

	head.Val, plow.Val = plow.Val, head.Val
	return plow
}
