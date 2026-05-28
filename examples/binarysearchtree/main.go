package main

import (
	"fmt"
	"github.com/divij2599/go-ds/binarysearchtree"
)

func main() {
	fmt.Println("Binary Search Tree Example")

	bst := binarysearchtree.New[int]()

	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	fmt.Println("InOrder:", bst.InOrder())
	fmt.Println("PreOrder:", bst.PreOrder())
	fmt.Println("PostOrder:", bst.PostOrder())

	fmt.Println("Search 40:", bst.Search(40))
	fmt.Println("Search 100:", bst.Search(100))

	minValue, ok := bst.Min()
	if ok {
		fmt.Println("Min:", minValue)
	}

	maxValue, ok := bst.Max()
	if ok {
		fmt.Println("Max:", maxValue)
	}

	fmt.Println("Height:", bst.Height())
	fmt.Println("Size:", bst.Size())

	bst.Delete(70)

	fmt.Println("After deleting 70:")
	fmt.Println("InOrder:", bst.InOrder())
	fmt.Println("Size:", bst.Size())

	fmt.Println()
	fmt.Println("String BST Example")

	stringBST := binarysearchtree.New[string]()

	stringBST.Insert("mango")
	stringBST.Insert("apple")
	stringBST.Insert("banana")
	stringBST.Insert("orange")

	fmt.Println("InOrder:", stringBST.InOrder())
	fmt.Println("Search banana:", stringBST.Search("banana"))
}
