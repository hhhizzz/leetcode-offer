package _24

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	next := head

	for next != nil {
		curr, next, next.Next = next, next.Next, curr
	}

	return curr
}
