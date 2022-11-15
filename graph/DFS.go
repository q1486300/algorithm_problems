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
	set := make(map[*Node]bool)
	stack.PushBack(node)
	set[node] = true
	fmt.Println(node.value)
	for stack.Len() != 0 {
		cur := stack.Back().Value.(*Node)
		stack.Remove(stack.Back())
		for _, next := range cur.nexts {
			_, ok := set[next]
			if !ok {
				stack.PushBack(cur)
				stack.PushBack(next)
				set[next] = true
				fmt.Println(next.value)
				break
			}
		}
	}
}
