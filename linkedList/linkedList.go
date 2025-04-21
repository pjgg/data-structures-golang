package linkedlist

import (
	"fmt"
)

// Nodo de la lista
type Node struct {
	data int
	next *Node
}

// Lista enlazada
type LinkedList struct {
	head *Node
}

// Insertar un nuevo nodo al final
func (l *LinkedList) Append(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// añadir un nodo por valor ' value < next'
func (l *LinkedList) Add(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next.data < value {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
}

// Eliminar un nodo por valor
func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	// Si el nodo a eliminar es el primero, entonces el head es el siguiente
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next // vamos a buscar el element cuyo siguiente elemento sea el que pensamos eliminar
	}

	if current.next != nil { // si no estmoa en la cola, entonces nos saltamos uno (que es el que eliminamos) y lo linkamos con el siguiente
		current.next = current.next.next
	}
}

// Mostrar la lista enlazada
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
