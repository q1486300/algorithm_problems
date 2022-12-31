package binary_tree

// 第一種方法
// 收集整棵樹的高度 h，和節點數 n
// 要求滿二元樹滿足: (2 ^ h) - 1 == n
func IsFull1(head *Node) bool {
	if head == nil {
		return true
	}
	allInfo := processIsFull1(head)
	return (1<<allInfo.height)-1 == allInfo.nodes
}

func processIsFull1(head *Node) IsFullInfo1 {
	if head == nil {
		return NewIsFullInfo(0, 0)
	}
	leftInfo := processIsFull1(head.left)
	rightInfo := processIsFull1(head.right)

	height := getMax(leftInfo.height, rightInfo.height) + 1
	nodes := leftInfo.nodes + rightInfo.nodes + 1
	return NewIsFullInfo(height, nodes)
}

type IsFullInfo1 struct {
	height, nodes int
}

func NewIsFullInfo(height, nodes int) IsFullInfo1 {
	return IsFullInfo1{
		height: height,
		nodes:  nodes,
	}
}

// 第二種方法
// 收集子樹是否是滿二元樹
// 收集子樹的高度
// 左樹滿 && 右樹滿 && 左右樹高度依樣 -> 整棵樹是滿的
func IsFull2(head *Node) bool {
	return processIsFull2(head).isFull
}

func processIsFull2(head *Node) IsFullInfo2 {
	if head == nil {
		return NewIsFullInfo2(true, 0)
	}
	leftInfo := processIsFull2(head.left)
	rightInfo := processIsFull2(head.right)

	isFull := leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height
	height := getMax(leftInfo.height, rightInfo.height) + 1
	return NewIsFullInfo2(isFull, height)
}

type IsFullInfo2 struct {
	isFull bool
	height int
}

func NewIsFullInfo2(isFull bool, height int) IsFullInfo2 {
	return IsFullInfo2{
		isFull: isFull,
		height: height,
	}
}
