package main

import "fmt"

func main() {
	p1 := &ListNode{Val: 1}
	p2 := &ListNode{Val: 2}
	p3 := &ListNode{Val: 3}
	p4 := &ListNode{Val: 4}
	p5 := &ListNode{Val: 5}
	p6 := &ListNode{Val: 6}

	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	p5.Next = p6
	p6.Next = p3
	fmt.Println(detectCycle(p1).Val)
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
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	p1, p2 := head.Next, head.Next.Next
	for p2 != nil {
		if p1 == p2 {
			break
		}
		if p2.Next == nil || p2.Next.Next == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	res := head
	for res != p2 {
		res = res.Next
		p2 = p2.Next
	}
	return res
}
