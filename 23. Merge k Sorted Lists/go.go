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
 func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		p, p1, p2 := res, lists[0], lists[1]
		for p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				p.Next = p1
				p = p.Next
                p1 = p1.Next
			} else {
				p.Next = p2
				p = p.Next
                p2 = p2.Next
			}
		}
		if p1 != nil {
			p.Next = p1
		}
		if p2 != nil {
			p.Next = p2
		}
		return res.Next
	}
	return mergeKLists([]*ListNode{mergeKLists(lists[0 : len(lists)/2]), mergeKLists(lists[len(lists)/2 : len(lists)])})
}