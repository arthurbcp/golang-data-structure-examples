package binary_tree

type BinaryTree[T any] struct {
	root *BinaryNode[T]
}

func (tree *BinaryTree[T]) Insert(key int, data T) {
	if tree.root == nil {
		tree.root = &BinaryNode[T]{key: key, data: data}
		return
	}
	tree.root.insert(key, data)
}

func (tree BinaryTree[T]) Search(key int) *BinaryNode[T] {
	return tree.root.search(key)
}

func (tree BinaryTree[T]) GetRoot() *BinaryNode[T] {
	return tree.root
}

func (tree BinaryTree[T]) GetMax() int {
	return tree.root.getMax()
}

func (tree BinaryTree[T]) GetMin() int {
	return tree.root.getMin()
}

func (tree BinaryTree[T]) getSize() int {
	return tree.root.getSize()
}
