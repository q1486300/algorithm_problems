package linkedlist

func ListPartition1(head *Node, pivot int) *Node {
	if head == nil {
		return head
	}
	cur := head
	i := 0
	for cur != nil {
		i++
		cur = cur.next
	}

	nodeArr := make([]*Node, i)
	cur = head
	for i := 0; i < len(nodeArr); i++ {
		nodeArr[i] = cur
		cur = cur.next
	}
	arrPartition(nodeArr, pivot)
	for i := 1; i < len(nodeArr); i++ {
		nodeArr[i-1].next = nodeArr[i]
	}
	nodeArr[len(nodeArr)-1].next = nil
	return nodeArr[0]
}

func arrPartition(nodeArr []*Node, pivot int) {
	small := -1
	big := len(nodeArr)
	index := 0
	for index < big {
		if nodeArr[index].value < pivot {
			small++
			swap(nodeArr, small, index)
			index++
		} else if nodeArr[index].value == pivot {
			index++
		} else {
			big--
			swap(nodeArr, big, index)
		}
	}
}

func swap(nodeArr []*Node, i, j int) {
	nodeArr[i], nodeArr[j] = nodeArr[j], nodeArr[i]
}

func ListPartition2(head *Node, pivot int) *Node {
	var (
		sH, sT *Node // small head, small tail
		eH, eT *Node // equal head, equal tail
		bH, bT *Node // big head, big tail
		next   *Node // save next node
	)
	// every node distributed to three lists
	for head != nil {
		next = head.next
		head.next = nil
		if head.value < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT.next = head
				sT = head
			}
		} else if head.value == pivot {
			if eH == nil {
				eH = head
				eT = head
			} else {
				eT.next = head
				eT = head
			}
		} else {
			if bH == nil {
				bH = head
				bT = head
			} else {
				bT.next = head
				bT = head
			}
		}
		head = next
	}
	// 小於區域的尾巴，連等於區域的頭; 等於區域的尾巴，連大於區域的頭
	if sT != nil { // 如果有小於區域
		sT.next = eH
		if eT == nil { // 沒有等於區域的話，就讓小於區域連大於區域的頭
			eT = sT
		}
	}

	// 有等於區域，eT -> 等於區域的尾巴
	// 有小於區域但是無等於區域，eT -> 小於區域的尾巴
	if eT != nil { // 小於區域跟等於區域至少有一部分
		eT.next = bH
	}

	if sH != nil {
		return sH
	}
	if eH != nil {
		return eH
	}
	return bH
}
