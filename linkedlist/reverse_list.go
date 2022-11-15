package linkedlist

// head
// a	->	b	->	c	->	null
// c	->	b	->	a	->	null
func ReverseLinkedList(head *Node) *Node {
	var pre, next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

func ReverseDoubleList(head *DoubleNode) *DoubleNode {
	var pre, next *DoubleNode
	for head != nil {
		next = head.next
		head.next = pre
		head.last = next
		pre = head
		head = next
	}
	return pre
}
