package binarytree

import (
	"fmt"
	"testing"
)

func LessThanComp(x interface{}, y interface{}) bool {
	return x.(int32) < y.(int32)
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


func TestBinaryTree_PrintValues(t *testing.T) {
	tree := New(LessThanComp)

	tree.Insert('F')
	tree.Insert('B')
	tree.Insert('G')
	tree.Insert('A')
	tree.Insert('D')
	tree.Insert('I')
	tree.Insert('C')
	tree.Insert('E')
	tree.Insert('H')
	fmt.Print("PostOrder: LRN - ")
	tree.PrintLRN()
	fmt.Println()

	fmt.Print("PreOrder: NLR - ")
	tree.PrintNLR()
	fmt.Println()

	fmt.Print("InOrder: LNR - ")
	tree.PrintLNR()
	fmt.Println()

	fmt.Print("OutOrder: RNL - ")
	tree.PrintRNL()
	fmt.Println()

	fmt.Print("Iterative Preorder (stack): NLR - ")
	tree.IterativePreOrder()
	fmt.Println()

	fmt.Print("Iterative InOrder (stack): LNR - ")
	tree.IterativeInOrder()
	fmt.Println()

	fmt.Print("Iterative BFS (queue): - ")
	tree.IterativeBFS()
	fmt.Println()

	fmt.Print("Morris Traversal (without queue or recursion): - ")
	tree.MorrisTraversal()
	fmt.Println()
}
