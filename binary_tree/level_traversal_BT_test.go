package binary_tree

import (
	"testing"
)

func TestLevelTraversal(t *testing.T) {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.right.left = NewNode(6)
	head.right.right = NewNode(7)

	LevelTraversal(head)
}
