package tree

// You're given a Node class that has a name
// and an array of optional children nodes.
// When put together, nodes form an acyclic
// tree-like structure.

// Implement the depthFirstSearch
// method on the Node class, which takes in
// an empty array, traverses the tree using the
// Depth-first Search approach (speciacally
// navigating the tree from left to right), stores
// all of the nodes' names in the input array,
// and returns it.

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}
