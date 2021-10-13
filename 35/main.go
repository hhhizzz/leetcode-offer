package _35

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	current := head
	// create new nodes in the old list
	for current != nil {
		current, current.Next = current.Next, &Node{Val: current.Val, Next: current.Next, Random: nil}
	}

	oldNode, newNode := head, head.Next

	// connect all the random node
	for {
		if oldNode.Random != nil {
			newNode.Random = oldNode.Random.Next
		}
		if newNode.Next == nil {
			break
		}
		oldNode, newNode = newNode.Next, newNode.Next.Next
	}

	// split all the new Node
	newHead := head.Next
	oldNode, newNode = head, head.Next
	for newNode.Next != nil {
		oldNode, oldNode.Next, newNode, newNode.Next = newNode.Next, newNode.Next, newNode.Next.Next, newNode.Next.Next
	}
	// after the iteration, the old node is still connected to the new last node, need to remove this link
	oldNode.Next = nil
	return newHead
}
