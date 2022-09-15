package main

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value, Size int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = NewBST(value)
			tree.Left.Size = 1
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = NewBST(value)
			tree.Right.Size = 1
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func (tree *BST) Contains(value int) bool {
	if value < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.Contains(value)
		}
	} else if value > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.Contains(value)
		}
	}
	return true
}

func (tree *BST) Remove(value int) *BST {
	return remove(tree, value)
}

func remove(tree *BST, value int) *BST {
	if tree == nil {
		return nil
	}
	if value < tree.Value {
		tree.Left = remove(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = remove(tree.Right, value)
	} else {
		if tree.Right == nil && tree.Left == nil {
			tree = nil
		} else if tree.Right == nil && tree.Left != nil {
			tree.Value = tree.Left.Value
			tree.Right = tree.Left.Right
			tree.Left = tree.Left.Left
		} else if tree.Left == nil && tree.Right != nil {
			tree.Value = tree.Right.Value
			tree.Left = tree.Right.Left
			tree.Right = tree.Right.Right
		} else {
			tree.Value = tree.Right.getMinNode().Value
			tree.Right = tree.Right.removeMin()
		}
	}

	if tree != nil {
		tree.Size = 1 + tree.Left.size() + tree.Right.size()
	}
	return tree
}

func (tree *BST) getMinNode() *BST {
	if tree.Left == nil {
		return tree
	}
	return tree.Left.getMinNode()
}

func (tree *BST) removeMin() *BST {
	if tree.Left == nil {
		return tree.Right
	}

	tree.Left = tree.Left.removeMin()
	tree.Size = 1 + tree.Left.size() + tree.Right.size()
	return tree
}

func (tree *BST) size() int {
	if tree == nil {
		return 0
	}
	return tree.Size
}
