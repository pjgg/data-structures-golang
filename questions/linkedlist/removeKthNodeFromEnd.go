package linkedlist

type LinkedList_removeKNodes struct {
	Value int
	Next  *LinkedList_removeKNodes
}

// Fast and Slow Pointer Technique

// remove element K from the end
func RemoveKthNodeFromEnd(head *LinkedList_removeKNodes, k int) {
	first := head
	second := head

	// Move second k steps ahead
	for i := 0; i < k; i++ {
		if second != nil {
			second = second.Next
		}
	}

	// If second is nil, we need to remove the head node
	if second == nil {
		// Copy the next node's data into head
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}

	// Move both pointers until second reaches the end
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the target node
	first.Next = first.Next.Next
}
