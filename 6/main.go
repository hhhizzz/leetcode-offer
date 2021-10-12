package _6

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result []int
	var current = head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	for i := 0; i < len(result)>>1; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}
