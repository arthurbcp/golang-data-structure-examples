package binary_tree

import "testing"

func TestGetData(t *testing.T) {
	node := &BinaryNode[string]{key: 1, data: "test"}
	if node.GetData() != "test" {
		t.Errorf("Search failed, expected: %s,  got: %s", "test", node.data)
	}
}
