package binarytree

type LessThanComparable func(a interface{}, b interface{}) bool

type BinaryTree struct {
	node     interface{}
	left     *BinaryTree
	right    *BinaryTree
	lessThan LessThanComparable
}

func New(comparable LessThanComparable) *BinaryTree {
	binaryTree := &BinaryTree{node: nil, left: nil, right: nil, lessThan: comparable}
	return binaryTree
}

func (tree *BinaryTree) Search(item interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}

	if tree.node == item {
		return tree
	}

	if tree.lessThan(item, tree.node) {
		// traverse left
		result := tree.left.Search(item)
		return result
	} else {
		// traverse right
		result := tree.right.Search(item)
		return result
	}
}

func (tree *BinaryTree) Insert(item interface{}) {
	if tree.node == nil {
		// create new
		tree.node = item
		tree.left = New(tree.lessThan)
		tree.right = New(tree.lessThan)
		return
	} else {
		// insert
		if tree.lessThan(item, tree.node) {
			tree.left.Insert(item)
		} else {
			tree.right.Insert(item)
		}
	}
}
