package linkedlist

// LeetCode 上的題目
// https://leetcode.com/problems/copy-list-with-random-pointer/
type NodeWithRandom struct {
	Val          int
	Next, Random *NodeWithRandom
}

func NewNodeWithRandom(val int) *NodeWithRandom {
	return &NodeWithRandom{
		Val: val,
	}
}

func CopyRandomList1(head *NodeWithRandom) *NodeWithRandom {
	// key 舊節點
	// value 新節點
	myMap := make(map[*NodeWithRandom]*NodeWithRandom)
	cur := head
	for cur != nil {
		myMap[cur] = NewNodeWithRandom(cur.Val)
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		// cur 舊
		// map[cur] 新
		// 新.next -> cur.next 複製節點找到
		myMap[cur].Next = myMap[cur.Next]
		myMap[cur].Random = myMap[cur.Random]
		cur = cur.Next
	}
	return myMap[head]
}

func CopyRandomList2(head *NodeWithRandom) *NodeWithRandom {
	if head == nil {
		return nil
	}
	cur := head
	var next *NodeWithRandom
	// 1 -> 2 -> 3 -> nil
	// 1 -> 1' -> 2 -> 2' -> 3 -> 3'
	for cur != nil {
		next = cur.Next
		cur.Next = NewNodeWithRandom(cur.Val)
		cur.Next.Next = next
		cur = next
	}

	cur = head
	var copyNode *NodeWithRandom
	// 1 1' 2 2' 3 3'
	// 依次設置 1' 2' 3' random指針
	for cur != nil {
		next = cur.Next.Next
		copyNode = cur.Next
		if cur.Random != nil {
			copyNode.Random = cur.Random.Next
		}
		cur = next
	}

	res := head.Next
	cur = head
	// 舊 新 混在一起，next方向上，random正確
	// next方向上，把新舊鏈結串列分離
	for cur != nil {
		next = cur.Next.Next
		copyNode = cur.Next
		cur.Next = next
		if next != nil {
			copyNode.Next = next.Next
		}
		cur = next
	}
	return res
}
