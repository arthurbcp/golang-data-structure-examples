package linked_list

import (
	"fmt"
	"testing"
)

var nodesArray = [10]int{100, 200, 53, 75, 713, 51, 633, 3, 644, 66}

func generateList() *LinkedList[string] {
	list := &LinkedList[string]{}
	for _, node := range nodesArray {
		list.Insert(node, "test "+fmt.Sprint(node))
	}
	return list
}

var list = generateList()

func TestInsert(t *testing.T) {
	currentNode := list.head
	i := 0
	for currentNode.next != nil {
		if currentNode.key != nodesArray[i] {
			t.Errorf("Insert failed, expected: %d,  got: %d", nodesArray[i], currentNode.key)
		}
		currentNode = currentNode.next
		i++
	}
}

func TestGetSize(t *testing.T) {
	emptyList := &LinkedList[string]{}
	size := emptyList.GetSize()
	if size != 0 {
		t.Errorf("GetSize failed, expected: %d,  got: %d", 0, size)
	}
	size = list.GetSize()
	if size != 10 {
		t.Errorf("GetSize failed, expected: %d,  got: %d", 10, size)
	}
}

func TestSearch(t *testing.T) {
	node1 := list.Search(244)
	if node1 != nil {
		t.Error("Search failed, expected: nil")
	}

	node2 := list.Search(75)
	if node2.data != "test 75" {
		t.Errorf("Search failed, expected: %s,  got: %s", "test 75", node2.data)
	}
}

func TestDelete(t *testing.T) {
	emptyList := &LinkedList[string]{}

	err := emptyList.Delete(1)
	if err == nil {
		t.Error("Delete failed, expected an error")
	}

	list.Delete(100)
	if list.head.key != 200 {
		t.Errorf("Delete failed, expected: %d,  got: %d", 200, list.head.key)
	}
	list.Delete(713)
	list.Delete(66)
	if list.GetSize() != 7 {
		t.Errorf("Delete failed, expected: %d,  got: %d", 8, list.GetSize())
	}
	err = list.Delete(999)
	if err == nil {
		t.Error("Delete failed, expected an error")
	}
}
