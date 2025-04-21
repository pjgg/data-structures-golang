package stack

import (
	"fmt"
)

// Nodo de la pila
type Node struct {
	data int
	next *Node
}

// Stack implementado con una lista enlazada
type Stack struct {
	head     *Node
	maxValue *Node
	minValue *Node
}

// Push agrega un nuevo elemento en la parte superior de la pila
func (s *Stack) Push(value int) {
	newNode := &Node{data: value, next: s.head}
	s.head = newNode

	// Actualizar el maxValue y minValue
	if s.maxValue == nil || value > s.maxValue.data {
		s.maxValue = newNode
	}
	if s.minValue == nil || value < s.minValue.data {
		s.minValue = newNode
	}
}

// Pop elimina y devuelve el elemento superior de la pila
func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	value := s.head.data
	s.head = s.head.next

	// Si eliminamos el maxValue o minValue, los recalculamos
	s.recalculateExtremes()

	return value, true
}

// Peek devuelve el elemento superior sin eliminarlo
func (s *Stack) Peek() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	return s.head.data, true
}

// Obtener el valor máximo actual de la pila
func (s *Stack) Max() (int, bool) {
	if s.maxValue == nil {
		return 0, false
	}
	return s.maxValue.data, true
}

// Obtener el valor mínimo actual de la pila
func (s *Stack) Min() (int, bool) {
	if s.minValue == nil {
		return 0, false
	}
	return s.minValue.data, true
}

// Método interno para recalcular los valores máximo y mínimo
func (s *Stack) recalculateExtremes() {
	current := s.head
	s.maxValue, s.minValue = current, current

	for current != nil {
		if s.maxValue == nil || current.data > s.maxValue.data {
			s.maxValue = current
		}
		if s.minValue == nil || current.data < s.minValue.data {
			s.minValue = current
		}
		current = current.next
	}
}

// Mostrar la pila
func (s *Stack) Print() {
	current := s.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
