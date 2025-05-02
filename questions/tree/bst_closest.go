package tree

// Write a function that takes in a Binary Search
// Tree (BST) and a target integer value and
// returns the closest value to that target value
// contained in the BST.
// You can assume that there will only be one
// closest value.

// Each BST node has an integer value , a
// left child node, and a right child node.
// A node is said to be a valid BST node if and
// only if it satis�es the BST property: its
// value is strictly greater than the values of
// every node to its left; its value is less than
// or equal to the values of every node to its

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return findClosestValue(tree, target, tree.Value)
}

func findClosestValue(node *BST, target int, closest int) int {
	if node == nil {
		return closest
	}

	if node.Value == target {
		return target
	}

	if abs(target-node.Value) < abs(target-closest) {
		closest = node.Value
	}

	if target < node.Value {
		return findClosestValue(node.Left, target, closest)
	}

	if target > node.Value {
		return findClosestValue(node.Right, target, closest)
	}

	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
