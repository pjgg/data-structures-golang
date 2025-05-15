package bst

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Head *Node
}

// Insertar ordenadamente
func (b *BST) Insert(value int) {
	b.Head = insert(b.Head, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	return node
}

// Buscar valor
func (b *BST) Search(value int) bool {
	return search(b.Head, value)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

// Eliminar valor
func (b *BST) Delete(value int) {
	b.Head = deleteNode(b.Head, value)
}

func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		// Nodo encontrado (que hay que borrar)

		// Caso 1: sin hijo izquierdo → reemplazar con derecho
		// Queremos eliminar el nodo 40 en este árbol:
		// 	50
		// 	/  \
		// 30  70
		// 	   /
		// 	  40
		// 		\
		// 		 45
		// Resultado del árbol tras eliminar 40
		// 	50
		// 	/  \
		// 30   70
		// 		/
		// 	  45

		if node.Left == nil && node.Right != nil {
			rightChild := node.Right
			node.Value = rightChild.Value
			node.Left = rightChild.Left
			node.Right = rightChild.Right
			return node
		}

		// Caso 2: sin hijo derecho → reemplazar con izquierdo
		if node.Right == nil && node.Left != nil {
			leftChild := node.Left
			node.Value = leftChild.Value
			node.Left = leftChild.Left
			node.Right = leftChild.Right
			return node
		}

		// Caso 3: tiene dos hijos
		// 	Eliminar nodo con 2 hijos:
		//      50
		//    /    \
		//  30      70
		//         /  \
		//       60    80
		// 		Eliminar 50 → reemplazo con 60 (mínimo del subárbol derecho)
		// → nuevo árbol:
		//          60
		//        /    \
		//      30      70
		//                 \
		//                 80
		// 👉 buscar el menor valor en el subárbol derecho
		if node.Right != nil && node.Left != nil {
			min := findMin(node.Right)
			// Reemplazar el valor del nodo con ese valor mínimo
			node.Value = min.Value
			// Eliminar el nodo duplicado (ya lo movimos arriba)
			node.Right = deleteNode(node.Right, min.Value)
		}
	}
	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = findMin(node.Left)
	}
	return node
}

// Recorridos
func (b *BST) InOrder() {
	inOrder(b.Head)
	fmt.Println()
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Print(node.Value, " ")
	inOrder(node.Right)
}
