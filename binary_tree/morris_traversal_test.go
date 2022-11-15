package binary_tree

import "testing"

func TestMorris(t *testing.T) {
	head := NewNode(4)
	head.left = NewNode(2)
	head.right = NewNode(6)
	head.left.left = NewNode(1)
	head.left.right = NewNode(3)
	head.right.left = NewNode(5)
	head.right.right = NewNode(7)

	PrintTree(head)

	MorrisPreOrder(head)
	MorrisInOrder(head)
	MorrisPostOrder(head)

	PrintTree(head)
}
