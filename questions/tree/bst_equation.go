package tree

// You're given a binary expression tree. Write a
// function to evaluate this tree mathematically
// and return a single resulting integer.
// All leaf nodes in the tree represent operands,
// which will always be positive integers. All of
// the other nodes represent operators. There
// are 4 operators supported, each of which is
// represented by a negative integer:
// • -1 : Addition operator, adding the left
// and right subtrees.
// • -2 : Subtraction operator, subtracting
// the right subtree from the left subtree.
// • -3 : Division operator, dividing the left
// subtree by the right subtree. If the

// (2*3) -2 + (8/3) = 6

type BinaryTreeEq struct {
	Value int

	Left  *BinaryTreeEq
	Right *BinaryTreeEq
}

func EvaluateExpressionTree(tree *BinaryTreeEq) int {

	//Base case: get to the leaf node
	if tree.Left == nil && tree.Right == nil {
		return tree.Value
	}

	leftValue := EvaluateExpressionTree(tree.Left)
	rightValue := EvaluateExpressionTree(tree.Right)

	if tree.Value == -1 {
		return leftValue + rightValue
	} else if tree.Value == -2 {
		return leftValue - rightValue
	} else if tree.Value == -3 {
		return leftValue / rightValue
	} else {
		return leftValue * rightValue
	}
}
