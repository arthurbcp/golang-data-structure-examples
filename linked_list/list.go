package linked_list

import "errors"

type LinkedList[T any] struct {
	head *Node[T]
}

func (list *LinkedList[T]) Insert(key int, data T) {
	if list.head == nil {
		list.head = &Node[T]{key: key, data: data}
		return
	}
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &Node[T]{key: key, data: data}
}

func (list LinkedList[T]) Search(key int) *Node[T] {
	currentNode := list.head
	for currentNode.key != key {
		currentNode = currentNode.next
		if currentNode == nil {
			return nil
		}
	}
	return currentNode
}

func (list *LinkedList[T]) Delete(key int) error {
	if list.GetSize() == 0 {
		return errors.New("you cannot delete a node from an empty list")
	}
	if list.head.key == key {
		list.head = list.head.next
		return nil
	}

	prevNode := list.head
	for prevNode.next.key != key {
		if prevNode.next.next == nil {
			return errors.New("you cannot delete a node that is not in the list")
		}
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
	return nil
}

func (list LinkedList[T]) GetSize() int {
	if list.head == nil {
		return 0
	}
	i := 1
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
		i++
	}
	return i
}
