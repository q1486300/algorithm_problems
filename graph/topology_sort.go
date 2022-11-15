package graph

import "container/list"

// directed graph and no loop
func SortedTopology(graph Graph) []*Node {
	// key: 某個節點，value: 剩餘的 incoming edges 數量
	inMap := make(map[*Node]int)
	// 只有 incoming edges 數量為 0 的節點，才進入這個隊列
	zeroInQueue := list.New()
	for _, node := range graph.nodes {
		inMap[node] = node.in
		if node.in == 0 {
			zeroInQueue.PushBack(node)
		}
	}

	var result []*Node
	for zeroInQueue.Len() != 0 {
		cur := zeroInQueue.Front().Value.(*Node)
		zeroInQueue.Remove(zeroInQueue.Front())
		result = append(result, cur)
		for _, next := range cur.nexts {
			inMap[next] = inMap[next] - 1
			if inMap[next] == 0 {
				zeroInQueue.PushBack(next)
			}
		}
	}
	return result
}
