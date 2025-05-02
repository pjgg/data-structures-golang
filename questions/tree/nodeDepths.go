package tree

// The distance between a node in a Binary Tree
// and the tree's root is called the node's depth.
// Write a function that takes in a Binary Tree
// and returns the sum of its nodes' depths.
// Each BinaryTree node has an integer
// value , a left child node, and a right
// child node. Children nodes can either be
// BinaryTree nodes themselves or None /
// null

type BTree struct {
	Value       int
	Left, Right *BTree
}

func NodeDepths(node *BTree) int {
	return nodeDepths(node, 0)
}

func nodeDepths(node *BTree, level int) int {
	if node == nil {
		return 0
	}

	// f(n,d) = l + f(nl, d + 1) + f(nr, d + 1)
	return level + nodeDepths(node.Left, level+1) + nodeDepths(node.Right, level+1)
}
