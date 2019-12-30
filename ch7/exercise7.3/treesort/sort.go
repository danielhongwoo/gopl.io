package treesort

import (
	"bytes"
	"strconv"
)

type Tree struct {
	value       int
	left, right *Tree
}

func (t *Tree) String() string {
	var writer bytes.Buffer
	t.traverse(&writer)
	return writer.String()
}

func (t *Tree) traverse(writer *bytes.Buffer) {
	if t == nil {
		return
	}
	t.left.traverse(writer)
	writer.Write([]byte(strconv.Itoa(t.value)))
	writer.Write([]byte(" "))
	t.right.traverse(writer)
}

func (root *Tree) Add(value int) *Tree {
	if root == nil {
		// Equivalent to return &Tree{value: value}.
		root = new(Tree)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = root.left.Add(value)
	} else {
		root.right = root.right.Add(value)
	}
	return root
}
