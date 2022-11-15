package linkedlist

import "math"

func GetIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

// 找到鏈結串列的第一個入環節點，如果無環返回 nil
func getLoopNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head.next
	fast := head.next.next
	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	// slow fast 相遇
	fast = head // fast -> walk again from head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

// 如果兩個鏈結串列都無環，返回第一個相交節點，如果不相交返回nil
func noLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1 := head1
	cur2 := head2
	n := 0
	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}
	if cur1 != cur2 {
		return nil
	}
	// n: 串列1長度 減去 串列2長度的值
	if n > 0 { // 誰長，就把它的 head 給 cur1
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	n = int(math.Abs(float64(n)))
	for n != 0 {
		n--
		cur1 = cur1.next
	}
	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

// 兩個有環鏈結串列，返回第一個相交節點，如果不相交返回nil
func bothLoop(head1, loop1, head2, loop2 *Node) *Node {
	var cur1, cur2 *Node
	if loop1 == loop2 {
		cur1 = head1
		cur2 = head2
		n := 0
		for cur1 != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.next
		}
		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}
		n = int(math.Abs(float64(n)))
		for n != 0 {
			n--
			cur1 = cur1.next
		}
		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}
