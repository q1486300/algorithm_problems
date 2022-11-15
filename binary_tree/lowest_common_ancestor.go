package binary_tree

func LowestCommonAncestor1(head, o1, o2 *Node) *Node {
	if head == nil {
		return nil
	}
	// key 的父節點是 value
	parentMap := make(map[*Node]*Node)
	parentMap[head] = nil
	fillParentMap(head, parentMap)

	o1Set := make(map[*Node]bool)
	cur := o1
	for parentMap[cur] != nil {
		o1Set[cur] = true
		cur = parentMap[cur]
	}
	o1Set[head] = true

	cur = o2
	for !o1Set[cur] {
		cur = parentMap[cur]
	}
	return cur
}

func fillParentMap(head *Node, parentMap map[*Node]*Node) {
	if head.left != nil {
		parentMap[head.left] = head
		fillParentMap(head.left, parentMap)
	}
	if head.right != nil {
		parentMap[head.right] = head
		fillParentMap(head.right, parentMap)
	}
}

func LowestCommonAncestor2(head, a, b *Node) *Node {
	return processLCA(head, a, b).ans
}

func processLCA(head, a, b *Node) LCAInfo {
	if head == nil {
		return NewLCAInfo(false, false, nil)
	}
	leftInfo := processLCA(head.left, a, b)
	rightInfo := processLCA(head.right, a, b)

	var findA, findB bool
	if head == a || leftInfo.findA || rightInfo.findA {
		findA = true
	}
	if head == b || leftInfo.findB || rightInfo.findB {
		findB = true
	}

	var ans *Node
	if leftInfo.ans != nil {
		ans = leftInfo.ans
	} else if rightInfo.ans != nil {
		ans = rightInfo.ans
	} else {
		if findA && findB {
			ans = head
		}
	}
	return NewLCAInfo(findA, findB, ans)
}

type LCAInfo struct {
	findA, findB bool
	ans          *Node
}

func NewLCAInfo(findA, findB bool, ans *Node) LCAInfo {
	return LCAInfo{
		findA: findA,
		findB: findB,
		ans:   ans,
	}
}
