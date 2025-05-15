package linkedlist

// build a double linked list

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) SetHead(node *Node) {
	if node == nil {
		return
	}
	if l.Head == node {
		return
	}
	l.Remove(node)

	node.Prev = nil
	node.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = node
	}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *DoublyLinkedList) SetTail(node *Node) {
	if node == nil {
		return
	}
	if l.Tail == node {
		return
	}
	l.Remove(node)

	node.Next = nil
	node.Prev = l.Tail
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
	if l.Head == nil {
		l.Head = node
	}
}

func (l *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if node == nil || nodeToInsert == nil || node == nodeToInsert {
		return
	}
	l.Remove(nodeToInsert)

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev != nil {
		node.Prev.Next = nodeToInsert
	} else {
		l.Head = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (l *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if node == nil || nodeToInsert == nil || node == nodeToInsert {
		return
	}
	l.Remove(nodeToInsert)

	nodeToInsert.Next = node.Next
	nodeToInsert.Prev = node

	if node.Next != nil {
		node.Next.Prev = nodeToInsert
	} else {
		l.Tail = nodeToInsert
	}
	node.Next = nodeToInsert
}

func (l *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 0 {
		l.SetHead(nodeToInsert)
		return
	}

	current := l.Head
	currentPosition := 1

	for current != nil && currentPosition != position {
		current = current.Next
		currentPosition++
	}

	if current != nil {
		l.InsertBefore(current, nodeToInsert)
	} else {
		l.SetTail(nodeToInsert)
	}
}

func (l *DoublyLinkedList) RemoveNodesWithValue(value int) {
	current := l.Head
	for current != nil {
		tmp := current
		current = current.Next
		if tmp.Value == value {
			l.Remove(tmp)
		}
	}
}

func (l *DoublyLinkedList) Remove(node *Node) {
	if node == nil {
		return
	}

	if node == l.Head {
		l.Head = l.Head.Next
	}
	if node == l.Tail {
		l.Tail = l.Tail.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}

func (l *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	current := l.Head
	for current != nil && current.Value != value {
		current = current.Next
	}
	return current != nil
}
