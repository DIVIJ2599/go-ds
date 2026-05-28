package binarysearchtree

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Node[T Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinarySearchTree[T Ordered] struct {
	Root *Node[T]
	size int
}

func New[T Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) Insert(value T) {
	if !t.Search(value) {
		t.Root = insertNode(t.Root, value)
		t.size++
	}
}

func insertNode[T Ordered](root *Node[T], value T) *Node[T] {
	if root == nil {
		return &Node[T]{Value: value}
	}

	if value < root.Value {
		root.Left = insertNode(root.Left, value)
	} else if value > root.Value {
		root.Right = insertNode(root.Right, value)
	}

	return root
}

func (t *BinarySearchTree[T]) Search(value T) bool {
	return searchNode(t.Root, value)
}

func searchNode[T Ordered](root *Node[T], value T) bool {
	if root == nil {
		return false
	}

	if root.Value == value {
		return true
	}

	if value < root.Value {
		return searchNode(root.Left, value)
	}

	return searchNode(root.Right, value)
}

func (t *BinarySearchTree[T]) Delete(value T) {
	if t.Search(value) {
		t.Root = deleteNode(t.Root, value)
		t.size--
	}
}

func deleteNode[T Ordered](root *Node[T], value T) *Node[T] {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		minNode := findMinNode(root.Right)
		root.Value = minNode.Value
		root.Right = deleteNode(root.Right, minNode.Value)
	}

	return root
}

func (t *BinarySearchTree[T]) Min() (T, bool) {
	var zero T

	if t.Root == nil {
		return zero, false
	}

	minNode := findMinNode(t.Root)

	return minNode.Value, true
}

func findMinNode[T Ordered](root *Node[T]) *Node[T] {
	currentNode := root

	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}

func (t *BinarySearchTree[T]) Max() (T, bool) {
	var zero T

	if t.Root == nil {
		return zero, false
	}

	currentNode := t.Root

	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}

	return currentNode.Value, true
}

func (t *BinarySearchTree[T]) InOrder() []T {
	values := []T{}
	inOrderHelper(t.Root, &values)

	return values
}

func inOrderHelper[T Ordered](root *Node[T], values *[]T) {
	if root == nil {
		return
	}

	inOrderHelper(root.Left, values)
	*values = append(*values, root.Value)
	inOrderHelper(root.Right, values)
}

func (t *BinarySearchTree[T]) PreOrder() []T {
	values := []T{}
	preOrderHelper(t.Root, &values)

	return values
}

func preOrderHelper[T Ordered](root *Node[T], values *[]T) {
	if root == nil {
		return
	}

	*values = append(*values, root.Value)
	preOrderHelper(root.Left, values)
	preOrderHelper(root.Right, values)
}

func (t *BinarySearchTree[T]) PostOrder() []T {
	values := []T{}
	postOrderHelper(t.Root, &values)

	return values
}

func postOrderHelper[T Ordered](root *Node[T], values *[]T) {
	if root == nil {
		return
	}

	postOrderHelper(root.Left, values)
	postOrderHelper(root.Right, values)
	*values = append(*values, root.Value)
}

func (t *BinarySearchTree[T]) Height() int {
	return heightHelper(t.Root)
}

func heightHelper[T Ordered](root *Node[T]) int {
	if root == nil {
		return -1
	}

	leftHeight := heightHelper(root.Left)
	rightHeight := heightHelper(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.Root == nil
}

func (t *BinarySearchTree[T]) Size() int {
	return t.size
}

func (t *BinarySearchTree[T]) Clear() {
	t.Root = nil
	t.size = 0
}
