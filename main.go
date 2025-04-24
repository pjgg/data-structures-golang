package main

import (
	queue "data-structures-golang-com/m/v2/Queue"
	"data-structures-golang-com/m/v2/bst"
	"data-structures-golang-com/m/v2/graph"
	hashmap "data-structures-golang-com/m/v2/hashMap"
	linkedlist "data-structures-golang-com/m/v2/linkedList"
	"data-structures-golang-com/m/v2/stack"
	"fmt"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Append(1)
	//list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Add(2)

	fmt.Println("Lista inicial:")
	list.Print()

	list.Delete(2)
	fmt.Println("Después de eliminar 2:")
	list.Print()

	list.Delete(4)
	fmt.Println("Después de eliminar 4:")
	list.Print()

	fmt.Println("======== HashMaps =========")

	hashMap := hashmap.NewHashMap(10)

	hashMap.Insert("foo", 42)
	hashMap.Insert("bar", 84)

	if value, found := hashMap.Get("foo"); found {
		fmt.Println("Value for 'foo':", value)
	} else {
		fmt.Println("Key 'foo' not found")
	}

	fmt.Println("======== Stack =========")
	// Crear una nueva pila
	s := &stack.Stack{}

	// Agregar elementos a la pila
	s.Push(10)
	s.Push(20)
	s.Push(5)
	s.Push(30)

	// Imprimir la pila
	fmt.Println("Estado de la pila:")
	s.Print()

	// Obtener el valor máximo y mínimo
	max, _ := s.Max()
	min, _ := s.Min()
	fmt.Println("Valor máximo:", max)
	fmt.Println("Valor mínimo:", min)

	// Peek del elemento superior
	top, _ := s.Peek()
	fmt.Println("Elemento en la cima:", top)

	// Pop del elemento superior
	popped, _ := s.Pop()
	fmt.Println("Elemento eliminado:", popped)

	// Imprimir nuevamente la pila
	fmt.Println("Estado de la pila después de Pop:")
	s.Print()

	fmt.Println("======== Queue =========")

	q := &queue.Queue{}

	// Agregar elementos a la cola
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(5)
	q.Enqueue(30)

	// Imprimir la cola
	fmt.Println("Estado de la cola:")
	q.Print()

	// Obtener el valor máximo y mínimo
	max, _ = q.Max()
	min, _ = q.Min()
	fmt.Println("Valor máximo:", max)
	fmt.Println("Valor mínimo:", min)

	// Peek del elemento frontal
	front, _ := q.Peek()
	fmt.Println("Elemento en el frente:", front)

	// Dequeue del elemento frontal
	dequeued, _ := q.Dequeue()
	fmt.Println("Elemento eliminado:", dequeued)

	// Imprimir nuevamente la cola
	fmt.Println("Estado de la cola después de Dequeue:")
	q.Print()

	fmt.Println("======== Grap =========")
	g := graph.NewGraph()

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	fmt.Println("DFS:")
	g.DFS(1, nil)

	fmt.Println("\nBFS:")
	g.BFS(1)

	fmt.Println("======== Trees / BST =========")
	tree := &bst.BST{}

	// Insertar elementos
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Print("InOrder traversal: ")
	tree.InOrder() // Debería imprimir: 20 30 40 50 60 70 80

	// Buscar elementos
	fmt.Println("Buscar 40:", tree.Search(40))   // true
	fmt.Println("Buscar 100:", tree.Search(100)) // false

	// Eliminar hoja
	tree.Delete(20)
	fmt.Print("Después de eliminar 20: ")
	tree.InOrder() // 30 40 50 60 70 80

	// Eliminar nodo con un hijo
	tree.Delete(30)
	fmt.Print("Después de eliminar 30: ")
	tree.InOrder() // 40 50 60 70 80

	// Eliminar nodo con dos hijos
	tree.Delete(50)
	fmt.Print("Después de eliminar 50: ")
	tree.InOrder() // 40 60 70 80
}
