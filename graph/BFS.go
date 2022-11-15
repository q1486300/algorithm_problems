package graph

import (
	"container/list"
	"fmt"
)

// 從 node 出發，進行廣度優先搜尋
func BFS(start *Node) {
	if start == nil {
		return
	}
	queue := list.New()
	set := make(map[*Node]bool)
	queue.PushBack(start)
	set[start] = true
	for queue.Len() != 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		fmt.Println(cur.value)
		for _, next := range cur.nexts {
			_, ok := set[next]
			if !ok {
				set[next] = true
				queue.PushBack(next)
			}
		}
	}
}
