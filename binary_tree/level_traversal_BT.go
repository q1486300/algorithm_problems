package binary_tree

import (
	"container/list"
	"fmt"
)

func LevelTraversal(head *Node) {
	if head == nil {
		return
	}
	queue := list.New()
	queue.PushBack(head)
	for queue.Len() != 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		fmt.Println(cur.value)
		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}
}
