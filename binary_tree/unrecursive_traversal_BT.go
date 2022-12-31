package binary_tree

import (
	"container/list"
	"fmt"
)

func unrecursivePreOrder(head *Node) {
	if head == nil {
		return
	}
	stack := list.New()
	stack.PushBack(head)
	for stack.Len() != 0 {
		head = stack.Back().Value.(*Node)
		stack.Remove(stack.Back())
		fmt.Printf("%d ", head.value)
		if head.right != nil {
			stack.PushBack(head.right)
		}
		if head.left != nil {
			stack.PushBack(head.left)
		}
	}
	fmt.Println()
}

func unrecursiveInOrder(head *Node) {
	if head == nil {
		return
	}
	stack := list.New()
	for stack.Len() != 0 || head != nil {
		if head != nil {
			stack.PushBack(head)
			head = head.left
		} else {
			head = stack.Back().Value.(*Node)
			stack.Remove(stack.Back())
			fmt.Printf("%d ", head.value)
			head = head.right
		}
	}
	fmt.Println()
}

func unrecursivePostOrder1(head *Node) {
	if head == nil {
		return
	}
	s1 := list.New()
	s1.PushBack(head)
	s2 := list.New()
	for s1.Len() != 0 {
		head = s1.Back().Value.(*Node) // 從 s1 堆疊彈出的順序: 頭 右 左
		s1.Remove(s1.Back())
		s2.PushBack(head)
		if head.left != nil {
			s1.PushBack(head.left)
		}
		if head.right != nil {
			s1.PushBack(head.right)
		}
	}
	// 從 s2 堆疊彈出的順序: 左 右 頭
	for s2.Len() != 0 {
		fmt.Printf("%d ", s2.Back().Value.(*Node).value)
		s2.Remove(s2.Back())
	}
	fmt.Println()
}

func unrecursivePostOrder2(head *Node) {
	if head == nil {
		return
	}
	stack := list.New()
	stack.PushBack(head)
	var c *Node
	for stack.Len() != 0 {
		c = stack.Back().Value.(*Node)
		if c.left != nil && c.left != head && c.right != head {
			stack.PushBack(c.left)
		} else if c.right != nil && c.right != head {
			stack.PushBack(c.right)
		} else {
			fmt.Printf("%d ", c.value)
			stack.Remove(stack.Back())
			head = c
		}
	}
	fmt.Println()
}
