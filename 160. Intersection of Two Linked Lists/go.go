package main

func main() {

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2, cnt1, cnt2 := headA, headB, 0, 0
	for p1 != nil {
		cnt1++
		p1 = p1.Next
	}

	for p2 != nil {
		cnt2++
		p2 = p2.Next
	}

	if cnt1 > cnt2 {
		for i := 0; i < cnt1-cnt2; i++ {
			headA = headA.Next
		}
	}
	if cnt1 < cnt2 {
		for i := 0; i < cnt2-cnt1; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
