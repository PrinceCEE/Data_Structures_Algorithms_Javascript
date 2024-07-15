package datastructures

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (bst *BST) Insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		bst.Root = &TreeNode{Value: value}
		return bst.Root
	}

	if value <= node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
			return node.Left
		} else {
			return bst.Insert(node.Left, value)
		}
	}

	if value > node.Value {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
			return node.Right
		} else {
			return bst.Insert(node.Right, value)
		}
	}

	return nil
}

func (bst *BST) Delete(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = bst.Delete(node.Left, value)
		return node
	}

	if value > node.Value {
		node.Right = bst.Delete(node.Right, value)
		return node
	}

	if node.Left == nil {
		return node.Right
	} else if node.Right == nil {
		return node.Left
	} else {
		successor := bst.FindMin(node)
		node.Value = successor.Value
		node.Right = bst.Delete(node.Right, successor.Value)
		return node
	}
}

func (bst *BST) FindMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = node.Left
	}
	return current
}

func (bst *BST) Search(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value == node.Value {
		return node
	}

	if value < node.Value {
		return bst.Search(node.Left, value)
	} else {
		return bst.Search(node.Right, value)
	}
}

// InOrder traversal follows left, root, right
func (bst *BST) TraverseInOrder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	bst.TraverseInOrder(node.Left, arr)
	*arr = append(*arr, node.Value)
	bst.TraverseInOrder(node.Right, arr)
}

// PreOrder traversal follows root, left, right
func (bst *BST) TraversePreOrder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	*arr = append(*arr, node.Value)
	bst.TraversePreOrder(node.Left, arr)
	bst.TraversePreOrder(node.Right, arr)
}

// PostOrder traversal follows left, right, root
func (bst *BST) TraversePostOrder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	bst.TraversePreOrder(node.Left, arr)
	bst.TraversePreOrder(node.Right, arr)
	*arr = append(*arr, node.Value)
}
