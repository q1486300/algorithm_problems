package binary_tree

import (
	"container/list"
)

func MaxWidthUseMap(head *Node) int {
	if head == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(head)
	// key 在 哪一層，value
	var levelMap = make(map[*Node]int)
	levelMap[head] = 1
	curLevel := 1      // 目前你正在統計哪一層的寬度
	curLevelNodes := 0 // 目前 curLevel 層，寬度是多少
	max := 0
	for queue.Len() != 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		curNodeLevel := levelMap[cur]
		if cur.left != nil {
			levelMap[cur.left] = curNodeLevel + 1
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			levelMap[cur.right] = curNodeLevel + 1
			queue.PushBack(cur.right)
		}
		if curNodeLevel == curLevel {
			curLevelNodes++
		} else {
			max = getMax(max, curLevelNodes)
			curLevel++
			curLevelNodes = 1
		}
	}
	max = getMax(max, curLevelNodes)
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxWidthNoMap(head *Node) int {
	if head == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(head)

	curEnd := head    // 目前層，最右節點是誰
	var nextEnd *Node // 下一層，最右節點是誰
	max := 0
	curLevelNodes := 0 // 目前層的節點數
	for queue.Len() != 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		if cur.left != nil {
			queue.PushBack(cur.left)
			nextEnd = cur.left
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
			nextEnd = cur.right
		}
		curLevelNodes++
		if cur == curEnd {
			max = getMax(max, curLevelNodes)
			curLevelNodes = 0
			curEnd = nextEnd
		}
	}
	return max
}
