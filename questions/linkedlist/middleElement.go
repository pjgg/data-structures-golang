package linkedlist

type LinkedList_mid struct {
	Value int
	Next  *LinkedList_mid
}

func MiddleNode(head *LinkedList_mid) *LinkedList_mid {
	current := head
	counter := 0

	for current != nil {
		counter++
		current = current.Next
	}

	mid := counter / 2
	current = head
	for i := 0; i < mid; i++ {
		current = current.Next
	}

	return current
}
