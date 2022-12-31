package binary_tree

func MaxDistance1(head *Node) int {
	return processMaxDistance(head).maxDistance
}

type MaxDistanceInfo struct {
	maxDistance int
	height      int
}

func NewMaxDistanceInfo(maxDistance, height int) MaxDistanceInfo {
	return MaxDistanceInfo{
		maxDistance: maxDistance,
		height:      height,
	}
}

func processMaxDistance(x *Node) MaxDistanceInfo {
	if x == nil {
		return NewMaxDistanceInfo(0, 0)
	}
	leftInfo := processMaxDistance(x.left)
	rightInfo := processMaxDistance(x.right)

	p1 := leftInfo.maxDistance
	p2 := rightInfo.maxDistance
	p3 := leftInfo.height + rightInfo.height + 1
	maxDistance := getMax(getMax(p1, p2), p3)
	height := getMax(leftInfo.height, rightInfo.height) + 1
	return NewMaxDistanceInfo(maxDistance, height)
}

func MaxDistance2(head *Node) int {
	if head == nil {
		return 0
	}
	arr := getPreList(head)
	parentMap := getParentMap(head)
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			max = getMax(max, distance(parentMap, arr[i], arr[j]))
		}
	}
	return max
}

func getPreList(head *Node) []*Node {
	var arr []*Node
	fillPrelist(head, &arr)
	return arr
}

func getParentMap(head *Node) map[*Node]*Node {
	myMap := make(map[*Node]*Node)
	myMap[head] = nil
	fillParentMap(head, myMap)
	return myMap
}

func distance(parentMap map[*Node]*Node, o1, o2 *Node) int {
	o1Set := make(map[*Node]bool)
	cur := o1
	o1Set[cur] = true
	for parentMap[cur] != nil {
		cur = parentMap[cur]
		o1Set[cur] = true
	}

	cur = o2
	for !o1Set[cur] {
		cur = parentMap[cur]
	}

	lowestAncestor := cur
	cur = o1
	distance1 := 1
	for cur != lowestAncestor {
		cur = parentMap[cur]
		distance1++
	}

	cur = o2
	distance2 := 1
	for cur != lowestAncestor {
		cur = parentMap[cur]
		distance2++
	}
	return distance1 + distance2 - 1
}
