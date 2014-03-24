package tree

import (
	"fmt"
)

type Tree struct {
	nodes map[string]*Tree
	value []byte
}

func (t *Tree) Add(k []string, v []byte) {
	if t.nodes == nil {
		t.nodes = make(map[string]*Tree)
	}
	if len(k) == 1 {
		t.nodes[k[0]] = &Tree{value: v}
		return
	}
	if _, ok := t.nodes[k[0]]; !ok {
		t.nodes[k[0]] = &Tree{}
	}
	t.nodes[k[0]].Add(k[1:], v)
}

func (t *Tree) Remove(k []string) {
	if len(k) == 1 {
		if t.nodes == nil {
			return
		}
		delete(t.nodes, k[0])
		return
	}
	if _, ok := t.nodes[k[0]]; !ok {
		return
	}
	t.nodes[k[0]].Remove(k[1:])
}

func (t *Tree) Get(k []string) []byte {
	if len(k) == 1 {
		if _, ok := t.nodes[k[0]]; !ok {
			return nil
		}
		return t.nodes[k[0]].value
	}
	if _, ok := t.nodes[k[0]]; !ok {
		return nil
	}
	return t.nodes[k[0]].Get(k[1:])
}

func (t *Tree) Print() {
	t.print(0, "*")
}

func (t *Tree) print(depth int, key string) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	if t.value == nil {
		fmt.Println(key)
	} else {
		fmt.Printf("%s: %s\n", key, string(t.value))
	}
	for k, v := range t.nodes {
		v.print(depth+1, k)
	}
}
