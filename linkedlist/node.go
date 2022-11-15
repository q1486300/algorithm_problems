package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(data int) *Node {
	return &Node{
		value: data,
	}
}

type DoubleNode struct {
	value      int
	last, next *DoubleNode
}

func NewDoubleNode(data int) *DoubleNode {
	return &DoubleNode{
		value: data,
	}
}

func generateLinkedListWithSlice(intArr []int) *Node {
	if intArr == nil {
		return nil
	}
	head := NewNode(intArr[0])
	pre := head
	for i := 1; i < len(intArr); i++ {
		cur := NewNode(intArr[i])
		pre.next = cur
		pre = cur
	}
	return head
}

func printLinkedList(node *Node) {
	fmt.Print("Linked List: ")
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Println()
}
