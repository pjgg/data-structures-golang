package tree

// Write a function that invert the BST.

type InvertBst struct {
	Value int

	Left  *InvertBst
	Right *InvertBst
}

func (node *InvertBst) InvertBinaryTree() {
	if node == nil {
		return
	}

	// Swap the left and right pointers
	node.Left, node.Right = node.Right, node.Left

	// Recursively invert the subtrees
	node.Left.InvertBinaryTree()
	node.Right.InvertBinaryTree()
}
