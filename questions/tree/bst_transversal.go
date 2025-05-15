package tree

type BST_trasversal struct {
	Value int

	Left  *BST_trasversal
	Right *BST_trasversal
}

func (tree *BST_trasversal) InOrderTraverse(array []int) []int {
	return inOrderTraverse(tree, array)
}

func inOrderTraverse(node *BST_trasversal, array []int) []int {
	if node == nil {
		return array
	}

	array = inOrderTraverse(node.Left, array)
	array = append(array, node.Value)
	array = inOrderTraverse(node.Right, array)
	return array
}

func (tree *BST_trasversal) PreOrderTraverse(array []int) []int {
	return preOrderTraverse(tree, array)
}

func preOrderTraverse(node *BST_trasversal, array []int) []int {
	if node == nil {
		return array
	}
	array = append(array, node.Value)
	array = preOrderTraverse(node.Left, array)
	array = preOrderTraverse(node.Right, array)
	return array
}

func (tree *BST_trasversal) PostOrderTraverse(array []int) []int {
	return postOrderTraverse(tree, array)
}

func postOrderTraverse(node *BST_trasversal, array []int) []int {
	if node == nil {
		return array
	}

	array = postOrderTraverse(node.Left, array)
	array = postOrderTraverse(node.Right, array)
	array = append(array, node.Value)

	return array
}
