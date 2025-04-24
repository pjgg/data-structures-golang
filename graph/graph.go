package graph

import (
	queue "data-structures-golang-com/m/v2/Queue"
	"fmt"
)

// Nodo (vértice del grafo (vertex))
type Node struct {
	value    int
	adjacent map[int]*Node // nodos adyacentes (edges)
}

// Grafo
type Graph struct {
	nodes map[int]*Node
}

// Crear un nuevo grafo
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}

// Añadir un nodo
func (g *Graph) AddNode(value int) {
	if _, exists := g.nodes[value]; !exists {
		g.nodes[value] = &Node{
			value:    value,
			adjacent: make(map[int]*Node),
		}
	}
}

// Eliminar un nodo (y sus aristas)
func (g *Graph) RemoveNode(value int) {
	if node, exists := g.nodes[value]; exists {
		// Eliminar el nodo de las listas de adyacencia de otros
		for _, neighbor := range node.adjacent {
			delete(neighbor.adjacent, value) // elimina el edge que va del vecino al nodo padre (porque vamos a borrar el nodo padre)
		}
		delete(g.nodes, value)
	}
}

// Añadir una arista (bidireccional)
func (g *Graph) AddEdge(value1, value2 int) {
	n1, ok1 := g.nodes[value1]
	n2, ok2 := g.nodes[value2]

	if ok1 && ok2 && value1 != value2 {
		n1.adjacent[value2] = n2
		n2.adjacent[value1] = n1
	}
}

// Eliminar una arista
func (g *Graph) RemoveEdge(value1, value2 int) {
	n1, ok1 := g.nodes[value1]
	n2, ok2 := g.nodes[value2]

	if ok1 && ok2 {
		delete(n1.adjacent, value2)
		delete(n2.adjacent, value1)
	}
}

// DFS - Depth First Search
func (g *Graph) DFS(start int, visited map[int]bool) {
	if visited == nil {
		visited = make(map[int]bool)
	}
	if _, exists := g.nodes[start]; !exists {
		return
	}
	if visited[start] {
		return
	}

	fmt.Print(start, " ")
	visited[start] = true

	for neighborID := range g.nodes[start].adjacent {
		g.DFS(neighborID, visited)
	}
}

// BFS - Breadth First Search
func (g *Graph) BFS(data int) {
	if _, exists := g.nodes[data]; !exists {
		return
	}

	visited := make(map[int]bool)
	q := &queue.Queue{}

	q.Enqueue(data)
	visited[data] = true

	for {
		start, ok := q.Dequeue()
		if !ok {
			break // Cola vacía
		}

		fmt.Print(start, " ")

		for neighborID := range g.nodes[start].adjacent {
			if !visited[neighborID] {
				visited[neighborID] = true
				q.Enqueue(neighborID)
			}
		}
	}
}
