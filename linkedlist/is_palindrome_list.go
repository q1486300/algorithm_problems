package linkedlist

import "container/list"

// need n extra space
func IsPalindrome1(head *Node) bool {
	stack := list.New()
	cur := head
	for cur != nil {
		stack.PushBack(cur)
		cur = cur.next
	}
	for head != nil {
		if head.value != stack.Back().Value.(*Node).value {
			return false
		}
		stack.Remove(stack.Back())
		head = head.next
	}
	return true
}

// need n/2 extra space
func IsPalindrome2(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	right := head.next
	cur := head
	for cur.next != nil && cur.next.next != nil {
		right = right.next
		cur = cur.next.next
	}

	stack := list.New()
	for right != nil {
		stack.PushBack(right)
		right = right.next
	}

	for stack.Len() != 0 {
		if head.value != stack.Back().Value.(*Node).value {
			return false
		}
		stack.Remove(stack.Back())
		head = head.next
	}
	return true
}

// need O(1) extra space
func IsPalindrome3(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	n1, n2 := head, head
	for n2.next != nil && n2.next.next != nil { // find mid node
		n1 = n1.next      // n1 -> mid
		n2 = n2.next.next // n2 -> end
	}
	// n1 中點

	n2 = n1.next // n2 -> right part first node
	n1.next = nil
	var n3 *Node
	for n2 != nil { // right part convert
		n3 = n2.next // n3 -> save next node
		n2.next = n1 // next of right node convert
		n1 = n2      // n1 move
		n2 = n3      // n2 move
	}

	n3 = n1   // n3 -> save last node
	n2 = head // n2 -> left first node
	res := true
	for n1 != nil && n2 != nil { // check palindrome
		if n1.value != n2.value {
			res = false
			break
		}
		n1 = n1.next // right to mid
		n2 = n2.next // left to mid
	}

	n1 = n3.next
	n3.next = nil
	for n1 != nil { // recover list
		n2 = n1.next
		n1.next = n3
		n3 = n1
		n1 = n2
	}
	return res
}
