package tree

// Write a BST class for a Binary Search Tree.
// The class should support:
// • Inserting values with the insert
// method.
// • Removing values with the remove
// method; this method should only
// remove the �rst instance of a given
// value.
// • Searching for values with the
// contains method.
// Note that you can't remove values from a
// single-node tree. In other words, calling the
// remove method on a single-node tree
// should simply not do anything.

type Tree struct {
	Value int

	Left  *Tree
	Right *Tree
}

func (tree *Tree) Insert(value int) *Tree {
	insert(tree, value)
	return tree
}

func insert(tree *Tree, value int) *Tree {
	if tree == nil {
		tree = &Tree{Value: value}
		return tree
	}

	if value < tree.Value {
		tree.Left = insert(tree.Left, value)
	} else {
		tree.Right = insert(tree.Right, value)
	}

	return tree
}

func (tree *Tree) Contains(value int) bool {
	return contains(tree, value)
}

func contains(tree *Tree, value int) bool {
	if tree == nil {
		return false
	}

	if value == tree.Value {
		return true
	} else if value < tree.Value {
		return contains(tree.Left, value)
	} else {
		return contains(tree.Right, value)
	}
}

func (tree *Tree) Remove(value int) *Tree {
	if tree.Left == nil && tree.Right == nil && tree.Value == value {
		// Return self unmodified
		return nil
	}

	return remove(tree, value)
}

func remove(tree *Tree, value int) *Tree {
	if tree == nil {
		return tree
	}

	if value < tree.Value {
		tree.Left = remove(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = remove(tree.Right, value)
	} else {
		// Node to be removed found
		if tree.Left == nil && tree.Right == nil {
			return nil
		}

		if tree.Left == nil && tree.Right != nil {
			rightChild := tree.Right
			tree.Value = rightChild.Value
			tree.Left = rightChild.Left
			tree.Right = rightChild.Right
			return tree
		}
		if tree.Right == nil && tree.Left != nil {
			leftChild := tree.Left
			tree.Value = leftChild.Value
			tree.Left = leftChild.Left
			tree.Right = leftChild.Right
			return tree
		}

		if tree.Right != nil && tree.Left != nil {
			// Node has two children: replace with smallest value in right subtree
			min := findMin(tree.Right)
			tree.Value = min.Value
			tree.Right = remove(tree.Right, min.Value)
		}
	}
	return tree
}

func findMin(node *Tree) *Tree {
	if node.Left != nil {
		node = findMin(node.Left)
	}

	return node
}
