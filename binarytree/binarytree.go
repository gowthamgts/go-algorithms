package binarytree

import (
	"fmt"
	"github.com/gowthamgts/go-algorithms/queue"
	"github.com/gowthamgts/go-algorithms/stack"
)

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

func (tree *BinaryTree) PrintLRN() {

	if tree.left != nil {
		tree.left.PrintLRN()
	}

	if tree.right != nil {
		tree.right.PrintLRN()
	}

	if tree.node != nil {
		fmt.Print(string(tree.node.(int32)))
		return
	}
}

func (tree *BinaryTree) PrintNLR() {
	if tree.node != nil {
		fmt.Print(string(tree.node.(int32)))
	}

	if tree.left != nil {
		tree.left.PrintNLR()
	}

	if tree.right != nil {
		tree.right.PrintNLR()
	}

	return
}

func (tree *BinaryTree) PrintLNR() {
	if tree.left != nil {
		tree.left.PrintLNR()
	}

	if tree.node != nil {
		fmt.Print(string(tree.node.(int32)))
	}

	if tree.right != nil {
		tree.right.PrintLNR()
	}

	return
}

func (tree *BinaryTree) PrintRNL() {

	if tree.right != nil {
		tree.right.PrintRNL()
	}

	if tree.node != nil {
		fmt.Print(string(tree.node.(int32)))
	}

	if tree.left != nil {
		tree.left.PrintRNL()
	}

	return
}

func (tree *BinaryTree) IterativePreOrder() {

	s := stack.New()
	s.Push(tree)

	for !s.IsEmpty() {

		tNode := s.Pop().(*BinaryTree)

		if tNode.node != nil {
			fmt.Print(string(tNode.node.(int32)))
		}

		if tNode.right != nil {
			s.Push(tNode.right)
		}

		if tNode.left != nil {
			s.Push(tNode.left)
		}
	}
}

func (tree *BinaryTree) IterativeInOrder() {
	s := stack.New()

	for !s.IsEmpty() || tree != nil {
		if tree != nil {
			s.Push(tree)
			tree = tree.left
		} else {
			// pop and visit
			tree = s.Pop().(*BinaryTree)
			if tree.node != nil {
				fmt.Print(string(tree.node.(int32)))
			}
			tree = tree.right
		}
	}
}

func (tree *BinaryTree) IterativeBFS() {
	q := queue.New()

	q.Enqueue(tree)

	for !q.IsEmpty() {
		qNode := q.Dequeue().(*BinaryTree)
		if qNode.node != nil {
			fmt.Print(string(qNode.node.(int32)))
		}

		if qNode.left != nil {
			q.Enqueue(qNode.left)
		}

		if qNode.right != nil {
			q.Enqueue(qNode.right)
		}
	}
}

func (tree *BinaryTree) MorrisTraversal() {

	current := tree

	for current != nil {
		if current.left == nil {
			if current.node != nil {
				fmt.Print(string(current.node.(int32)))
			}
			current = current.right
		} else {
			leftRoot := current.left
			for leftRoot.right != nil && leftRoot.right != current {
				leftRoot = leftRoot.right
			}

			if leftRoot.right == nil {
				leftRoot.right = current
				current = current.left
			} else {
				leftRoot.right = nil
				if current.node != nil {
					fmt.Print(string(current.node.(int32)))
				}
				current = current.right
			}
		}
	}

}