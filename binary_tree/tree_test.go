package binary_tree

import (
	"fmt"
	"testing"
)

var nodesArray = [10]int{100, 200, 53, 75, 713, 51, 633, 3, 644, 66}

func generateTree() *BinaryTree[string] {
	tree := &BinaryTree[string]{}
	for _, node := range nodesArray {
		tree.Insert(node, "test "+fmt.Sprint(node))
	}
	return tree
}

var tree = generateTree()

func TestInsert(t *testing.T) {
	root := tree.root
	if root.key != 100 {
		t.Errorf("Insert failed, expected: %d,  got: %d", 100, root.key)
	}
	if root.data != "test 100" {
		t.Errorf("Insert failed, expected: %s,  got: %s", "test", root.data)
	}
	if root.left.key != 53 {
		t.Errorf("Insert failed, expected: %d,  got: %d", 53, root.left.key)
	}
	if root.right.key != 200 {
		t.Errorf("Insert failed, expected: %d,  got: %d", 200, root.right.key)
	}
}

func TestSearch(t *testing.T) {
	node := tree.Search(75)
	if node.data != "test 75" {
		t.Errorf("Search failed, expected: %s,  got: %s", "test 75", node.data)
	}
	if node.left.key != 66 {
		t.Errorf("Search failed, expected: %d,  got: %d", 66, node.left.key)
	}
	if node.right != nil {
		t.Errorf("Search failed, expected: nil")
	}
}

func TestGetSize(t *testing.T) {
	size := tree.getSize()
	if size != 10 {
		t.Errorf("Size failed, expected: %d,  got: %d", 10, size)
	}
}

func TestGetMax(t *testing.T) {
	max := tree.GetMax()
	if max != 713 {
		t.Errorf("Max failed, expected: %d,  got: %d", 713, max)
	}
}

func TestGetMin(t *testing.T) {
	min := tree.GetMin()
	if min != 3 {
		t.Errorf("Min failed, expected: %d,  got: %d", 3, min)
	}
}

func TestGetRoot(t *testing.T) {
	root := tree.GetRoot()
	if root.key != 100 {
		t.Errorf("GetRoot failed, expected: %d,  got: %d", 100, root.key)
	}
}
