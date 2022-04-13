package linked_list

type Node[T any] struct {
	key  int
	data T
	next *Node[T]
}
