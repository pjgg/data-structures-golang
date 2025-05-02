package tree

// Write a function that takes in a Binary Tree
// and returns a list of its branch sums ordered
// from leftmost branch sum to rightmost
// branch sum.
// A branch sum is the sum of all values in a
// Binary Tree branch. A Binary Tree branch is a
// path of nodes in a tree that starts at the root
// node and ends at any leaf node.
// Each BinaryTree node has an integer
// value , a left child node, and a right
// child node. Children nodes can either be
// BinaryTree nodes themselves or None /
// null

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var sums []int
	return calculateBranchSums(root, 0, sums)
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums []int) []int {
	if node == nil {
		return sums
	}

	runningSum = runningSum + node.Value

	// Then this is a leaf node
	if node.Left == nil && node.Right == nil {
		sums = append(sums, runningSum)
	}

	sums = calculateBranchSums(node.Left, runningSum, sums)
	sums = calculateBranchSums(node.Right, runningSum, sums)

	return sums
}
