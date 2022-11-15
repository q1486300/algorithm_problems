package binary_tree

import (
	"fmt"
	"testing"
)

func TestUnrecursiveTraversal(t *testing.T) {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.right.left = NewNode(6)
	head.right.right = NewNode(7)

	unrecursivePreOrder(head)
	fmt.Println("=============")
	unrecursiveInOrder(head)
	fmt.Println("=============")
	unrecursivePostOrder1(head)
	fmt.Println("=============")
	unrecursivePostOrder2(head)
	fmt.Println("=============")
}
