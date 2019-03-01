package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 7}

	print(addTwoNumbers(l1, l2))
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	v := l1.Val + l2.Val
	res := &ListNode{Val: v % 10}
	res.Next = addTwoNumbers(l1.Next, l2.Next)
	if v >= 10 {
		res.Next = addTwoNumbers(&ListNode{Val: v / 10}, res.Next)
	}
	return res
}
