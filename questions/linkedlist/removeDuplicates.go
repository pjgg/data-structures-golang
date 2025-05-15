package linkedlist

type LinkedList_removeDuplicates struct {
	Value int
	Next  *LinkedList_removeDuplicates
}

func RemoveDuplicatesFromLinkedList(head *LinkedList_removeDuplicates) *LinkedList_removeDuplicates {
	if head == nil {
		return nil
	}

	current := head

	for current.Next != nil {
		if current.Value != current.Next.Value {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}

	return head
}
