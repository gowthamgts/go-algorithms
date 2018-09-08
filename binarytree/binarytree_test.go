package binarytree

import "testing"

func LessThanComp(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func TestBinaryTree_Insert(t *testing.T) {
	tree := New(LessThanComp)

	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(2)

	rootTree := tree.Search(3)

	if rootTree == nil {
		t.Error("Tree not found")
	}

	nilTree := tree.Search(100)

	if nilTree != nil {
		t.Error("Tree should not be present")
	}

	leftTree := tree.Search(1)
	if leftTree == nil {
		t.Error("Tree not found")
	}

	rightTree := tree.Search(2)
	if rightTree == nil {
		t.Error("Tree not found")
	}
}
