package queue

import (
	"fmt"
)

// Nodo de la cola
type Node struct {
	data int
	next *Node
	prev *Node
}

// Cola (FIFO) implementada con una lista doblemente enlazada
type Queue struct {
	head     *Node
	tail     *Node
	maxValue *Node
	minValue *Node
}

// Enqueue agrega un nuevo elemento al final de la cola
func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.tail == nil {
		q.head, q.tail = newNode, newNode
	} else {
		q.tail.next = newNode
		newNode.prev = q.tail
		q.tail = newNode
	}

	// Actualizar el maxValue y minValue
	if q.maxValue == nil || value > q.maxValue.data {
		q.maxValue = newNode
	}
	if q.minValue == nil || value < q.minValue.data {
		q.minValue = newNode
	}
}

// Dequeue elimina y devuelve el elemento al frente de la cola
func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	value := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	} else {
		q.head.prev = nil
	}

	// Si eliminamos el maxValue o minValue, los recalculamos
	q.recalculateExtremes()

	return value, true
}

// Peek devuelve el elemento al frente sin eliminarlo
func (q *Queue) Peek() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	return q.head.data, true
}

// Max devuelve el valor máximo actual en la cola
func (q *Queue) Max() (int, bool) {
	if q.maxValue == nil {
		return 0, false
	}
	return q.maxValue.data, true
}

// Min devuelve el valor mínimo actual en la cola
func (q *Queue) Min() (int, bool) {
	if q.minValue == nil {
		return 0, false
	}
	return q.minValue.data, true
}

// Método interno para recalcular los valores máximo y mínimo
func (q *Queue) recalculateExtremes() {
	current := q.head
	q.maxValue, q.minValue = current, current

	for current != nil {
		if q.maxValue == nil || current.data > q.maxValue.data {
			q.maxValue = current
		}
		if q.minValue == nil || current.data < q.minValue.data {
			q.minValue = current
		}
		current = current.next
	}
}

// Imprimir la cola
func (q *Queue) Print() {
	current := q.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}
