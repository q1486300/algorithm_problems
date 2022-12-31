package linkedlist

// 1(*), 2
// 1, 2(*), 3
// 1, 2(*), 3, 4
func MidOrUpMidNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	// 有三個以上的節點
	slow := head.next
	fast := head.next.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 1, 2(*)
// 1, 2(*), 3
// 1, 2, 3(*), 4
func MidOrDownMidNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	slow := head.next
	fast := head.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 1(*), 2, 3
// 1(*), 2, 3, 4
// 1, 2(*), 3, 4, 5
func MidOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head
	fast := head.next.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 1(*), 2
// 1(*), 2, 3
// 1, 2(*), 3, 4
// 1, 2(*), 3, 4, 5
func MidOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	slow := head
	fast := head.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
