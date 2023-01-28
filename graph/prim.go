package graph

import "container/heap"

func PrimMST(graph Graph) []*Edge {
	// 解鎖的邊進入最小堆積樹
	edgePQ := NewEdgePriorityQueue()
	// 哪些點被解鎖出來了
	nodeSet := make(map[*Node]struct{})
	// 依次挑選的邊進入 result 中
	var result []*Edge

	for _, node := range graph.nodes { // 隨便挑了一個節點
		// node 是開始點
		_, ok := nodeSet[node]
		if !ok {
			nodeSet[node] = struct{}{}
			for _, edge := range node.edges { // 由一個節點，解鎖所有相鄰的邊
				heap.Push(edgePQ, edge)
			}
			for edgePQ.Len() != 0 {
				curEdge := heap.Pop(edgePQ).(*Edge) // 彈出解鎖的邊中，最小的邊
				toNode := curEdge.to
				_, ok = nodeSet[toNode]
				if !ok { // 集合中不包含 toNode 時，toNode 是新節點
					nodeSet[toNode] = struct{}{}
					result = append(result, curEdge)
					for _, nextEdge := range toNode.edges {
						heap.Push(edgePQ, nextEdge)
					}
				}
			}
		}
		// break	// 如果 graph 整體連通的話循環只會走一次，不完全連通(即森林)則會找到所有的最小生成樹
	}
	return result
}
