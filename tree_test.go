package tree

import (
	"bytes"
	"testing"
)

func TestTree(t *testing.T) {
	var tree Tree
	var k = []string{"foo", "bar"}
	var v = []byte("baz")

	vals := tree.Get(k)
	if vals != nil {
		t.Fatalf("unexpected value: %+v", vals)
	}

	tree.Remove(k) // must not panic

	tree.Add(k, v)
	vals = tree.Get(k)
	if bytes.Compare(vals, v) != 0 {
		t.Fatalf("unexpected value: %+v", vals)
	}

	// Add same k/v
	tree.Add(k, v)
	vals = tree.Get(k)
	if bytes.Compare(vals, v) != 0 {
		t.Fatalf("unexpected value: %+v", vals)
	}

	tree.Remove(k)
	vals = tree.Get(k)
	if vals != nil {
		t.Fatalf("unexpected value: %+v", vals)
	}
}
