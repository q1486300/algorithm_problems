package graph

import (
	"container/list"
	"fmt"
)

func DFS(node *Node) {
	if node == nil {
		return
	}
	stack := list.New()
	set := make(map[*Node]struct{})
	stack.PushBack(node)
	set[node] = struct{}{}
	fmt.Println(node.value) // 進堆疊時就印出(處理)
	for stack.Len() != 0 {
		cur := stack.Back().Value.(*Node)
		stack.Remove(stack.Back())
		for _, next := range cur.nexts {
			_, ok := set[next]
			if !ok {
				stack.PushBack(cur)
				stack.PushBack(next)
				set[next] = struct{}{}
				fmt.Println(next.value) // 進堆疊時就印出(處理)
				break
			}
		}
	}
}
