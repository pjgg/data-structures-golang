package tree

import "math"

// Validate BST
// Write a function that takes in a potentially
// invalid Binary Search Tree (BST) and returns a
// boolean representing whether the BST is
// valid.
// Each BST node has an integer value , a
// left child node, and a right child node.
// A node is said to be a valid BST node if and
// only if it satis�es the BST property: its
// value is strictly greater than the values of
// every node to its left; its value is less than
// or equal to the values of every node to its
// right; and its children nodes are either valid
// BST nodes themselves or None / null .
// A BST is valid if and only if all of its nodes are
// valid BST nodes

// Explanation
// Every node has a max and a min value. If we are evaluating the Left nodes then the Max value should be the current node value (node.Value),
// and if we evaluate the right branch then the minimum value is gonna be the current node value
type BST_validate struct {
	Value int

	Left  *BST_validate
	Right *BST_validate
}

func (tree *BST_validate) ValidateBst() bool {
	// Write your code here.
	return isValid(tree, math.MaxInt, math.MinInt)
}

func isValid(node *BST_validate, max, min int) bool {
	if node == nil {
		return true
	}

	if node.Value < min || node.Value >= max {
		return false
	}

	return isValid(node.Left, node.Value, min) && isValid(node.Right, max, node.Value)
}
