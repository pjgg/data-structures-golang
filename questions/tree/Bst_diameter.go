package tree

// Write a function that takes in a Binary Tree
// and returns its diameter. The diameter of a
// binary tree is de�ned as the length of its
// longest path, even if that path doesn't pass
// through the root of the tree.
// A path is a collection of connected nodes in a
// tree, where no node is connected to more
// than two other nodes. The length of a path is
// the number of edges between the path's �rst
// node and its last node.

// This is an input class. Do not edit.
type BinaryTree_diameter struct {
	Value int

	Left  *BinaryTree_diameter
	Right *BinaryTree_diameter
}

func BinaryTreeDiameter(root *BinaryTree_diameter) int {
	maxPath := 0
	DFT(root, &maxPath)
	return maxPath
}

func DFT(node *BinaryTree_diameter, maxPath *int) int {
	if node == nil {
		return 0
	}

	// leftAmountOfEdges and rightAmountOfEdges are the heights(amount of edges) of left and right subtrees.
	leftAmountOfEdges := DFT(node.Left, maxPath)
	rightAmountOfEdges := DFT(node.Right, maxPath)

	// The longest path through the current node is leftAmountOfEdges + rightAmountOfEdges.
	pathThroughNode := leftAmountOfEdges + rightAmountOfEdges

	if pathThroughNode > *maxPath {
		*maxPath = pathThroughNode
	}

	return max(leftAmountOfEdges, rightAmountOfEdges) + 1
}

func max(data1, data2 int) int {
	if data1 >= data2 {
		return data1
	}
	return data2
}
