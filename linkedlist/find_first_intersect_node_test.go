package linkedlist

import "testing"

func TestGetIntersectNode(t *testing.T) {
	var heads []*Node
	var nodeValues = [][]int{
		{1, 2, 3, 4, 5, 6, 7}, // head1 無環
		{0, 9, 8},             // head2 無環 (與 heads[0] 相交於 6)

		{1, 2, 3, 4, 5, 6, 7}, // head1 有環 (第一個入環節點 4)
		{0, 9, 8},             // head2 有環 (與 heads[2] 相交於 2)
		{0, 9, 8},             // head2 有環 (與 heads[2] 相交於 6)
	}
	generateHeads(&heads, nodeValues)

	heads[1].next.next.next = heads[0].next.next.next.next.next
	heads[2].next.next.next.next.next.next.next = heads[2].next.next.next
	heads[3].next.next.next = heads[2].next
	heads[4].next.next.next = heads[2].next.next.next.next.next

	t.Log(GetIntersectNode(heads[0], heads[1]))

	t.Log(GetIntersectNode(heads[2], heads[3]))

	t.Log(GetIntersectNode(heads[2], heads[4]))
}
