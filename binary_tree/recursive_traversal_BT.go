package binary_tree

import "fmt"

// 前序印出所有節點
func PreOrder(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	PreOrder(head.left)
	PreOrder(head.right)
}

func InOrder(head *Node) {
	if head == nil {
		return
	}
	InOrder(head.left)
	fmt.Println(head.value)
	InOrder(head.right)
}

func PostOrder(head *Node) {
	if head == nil {
		return
	}
	PostOrder(head.left)
	PostOrder(head.right)
	fmt.Println(head.value)
}
