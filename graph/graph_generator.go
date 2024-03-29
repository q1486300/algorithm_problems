package graph

// matrix 所有的邊
// N * 3 的矩陣
// [weight, from節點上的值, to節點上的值]
// [5, 0, 7]
// [3, 0, 1]
func createGraph(matrix [][]int) Graph {
	graph := NewGraph()
	for i := 0; i < len(matrix); i++ {
		// 拿到每一條邊，matrix[i]
		weight := matrix[i][0]
		from := matrix[i][1]
		to := matrix[i][2]

		fromNode, ok := graph.nodes[from]
		if !ok {
			graph.nodes[from] = NewNode(from)
			fromNode = graph.nodes[from]
		}

		toNode, ok := graph.nodes[to]
		if !ok {
			graph.nodes[to] = NewNode(to)
			toNode = graph.nodes[to]
		}

		newEdge := NewEdge(weight, fromNode, toNode)
		graph.edges[newEdge] = true

		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.edges = append(fromNode.edges, newEdge)
		fromNode.out++
		toNode.in++
	}
	return graph
}
