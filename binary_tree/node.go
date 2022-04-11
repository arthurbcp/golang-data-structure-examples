package binary_tree

type BinaryNode[T any] struct {
	key   int
	data  T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func (node *BinaryNode[T]) insert(key int, data T) {
	if key < node.key {
		if node.left == nil {
			node.left = &BinaryNode[T]{key: key, data: data}
		} else {
			node.left.insert(key, data)
		}
	} else if key > node.key {
		if node.right == nil {
			node.right = &BinaryNode[T]{key: key, data: data}
		} else {
			node.right.insert(key, data)
		}
	}
}

func (node *BinaryNode[T]) search(key int) *BinaryNode[T] {
	if key < node.key {
		return node.left.search(key)
	} else if key > node.key {
		return node.right.search(key)
	}

	return node
}

func (node *BinaryNode[T]) GetData() T {
	return node.data
}

func (node *BinaryNode[T]) getSize() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.getSize() + node.right.getSize()
}

func (node *BinaryNode[T]) getMax() int {
	if node.right == nil {
		return node.key
	}
	return node.right.getMax()
}

func (node *BinaryNode[T]) getMin() int {
	if node.left == nil {
		return node.key
	}
	return node.left.getMin()
}
