package binary_tree

func IsBST1(head *Node) bool {
	if head == nil {
		return true
	}
	var arr []*Node
	in(head, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i].value <= arr[i-1].value {
			return false
		}
	}
	return true
}

func in(head *Node, arr *[]*Node) {
	if head == nil {
		return
	}
	in(head.left, arr)
	*arr = append(*arr, head)
	in(head.right, arr)
}

func IsBST2(head *Node) bool {
	if head == nil {
		return true
	}
	return processIsBST(head).isBST
}

func processIsBST(head *Node) *IsBSTInfo {
	if head == nil {
		return nil
	}
	leftInfo := processIsBST(head.left)
	rightInfo := processIsBST(head.right)

	min := head.value
	max := head.value
	if leftInfo != nil {
		min = getMin(min, leftInfo.min)
		max = getMax(max, leftInfo.max)
	}
	if rightInfo != nil {
		min = getMin(min, rightInfo.min)
		max = getMax(max, rightInfo.max)
	}

	isBST := true
	if leftInfo != nil && (!leftInfo.isBST || leftInfo.max >= head.value) {
		isBST = false
	}
	if rightInfo != nil && (!rightInfo.isBST || rightInfo.min <= head.value) {
		isBST = false
	}
	return NewIsBSTInfo(isBST, max, min)
}

type IsBSTInfo struct {
	isBST    bool
	max, min int
}

func NewIsBSTInfo(i bool, max, min int) *IsBSTInfo {
	return &IsBSTInfo{
		isBST: i,
		max:   max,
		min:   min,
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
