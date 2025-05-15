package tree

// given a sorted array with distinct elements, develop a function that construct a BST with the lowest depth

func MinHeightBST(array []int) *BST_smallest {
	if len(array) == 0 {
		return nil
	}

	mid := len(array) / 2
	bst := new(BST_smallest)
	bst.Value = array[mid]
	bst.Left = MinHeightBST(array[0:mid])
	bst.Right = MinHeightBST(array[mid+1 : len(array)])
	return bst
}

type BST_smallest struct {
	Value int

	Left  *BST_smallest
	Right *BST_smallest
}
