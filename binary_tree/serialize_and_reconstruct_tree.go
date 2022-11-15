package binary_tree

import (
	"container/list"
	"strconv"
)

/*
二元樹可以透過 前序、後序，或按層遍歷的方式序列化和反序列化
但是，二元樹無法透過中序遍歷的方式實現序列化和反序列化
因為不同的兩棵樹，可能得到同樣的中序列表，即使補了空位置也可能一樣
例如下列兩棵樹
		__2
       /
      1
      和
      1__
         \
          2
補足空位置的中序遍歷結果都是 { nil, 1, nil, 2, nil }
*/

func PreSerial(head *Node) *list.List {
	ans := list.New()
	pres(head, ans)
	return ans
}

func pres(head *Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		ans.PushBack(strconv.Itoa(head.value))
		pres(head.left, ans)
		pres(head.right, ans)
	}
}

func InSerial(head *Node) *list.List {
	ans := list.New()
	ins(head, ans)
	return ans
}

func ins(head *Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		ins(head.left, ans)
		ans.PushBack(strconv.Itoa(head.value))
		ins(head.right, ans)
	}
}

func PostSerial(head *Node) *list.List {
	ans := list.New()
	posts(head, ans)
	return ans
}

func posts(head *Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		posts(head.left, ans)
		posts(head.right, ans)
		ans.PushBack(strconv.Itoa(head.value))
	}
}

func BuildByPreQueue(preList *list.List) *Node {
	if preList == nil || preList.Len() == 0 {
		return nil
	}
	return preb(preList)
}

func preb(preList *list.List) *Node {
	value, ok := preList.Front().Value.(string)
	preList.Remove(preList.Front())
	if !ok {
		return nil
	}
	v, _ := strconv.Atoi(value)
	head := NewNode(v)
	head.left = preb(preList)
	head.right = preb(preList)
	return head
}

func BuildByPostQueue(postList *list.List) *Node {
	if postList == nil || postList.Len() == 0 {
		return nil
	}
	// 左右中 -> stack (中右左)
	stack := list.New()
	for postList.Len() != 0 {
		stack.PushBack(postList.Front().Value)
		postList.Remove(postList.Front())
	}
	return postb(stack)
}

func postb(postStack *list.List) *Node {
	value, ok := postStack.Back().Value.(string)
	postStack.Remove(postStack.Back())
	if !ok {
		return nil
	}
	v, _ := strconv.Atoi(value)
	head := NewNode(v)
	head.right = postb(postStack)
	head.left = postb(postStack)
	return head
}

func LevelSerial(head *Node) *list.List {
	ans := list.New()
	if head == nil {
		ans.PushBack(nil)
	} else {
		ans.PushBack(strconv.Itoa(head.value))
		queue := list.New()
		queue.PushBack(head)
		for queue.Len() != 0 {
			head := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())
			if head.left != nil {
				ans.PushBack(strconv.Itoa(head.left.value))
				queue.PushBack(head.left)
			} else {
				ans.PushBack(nil)
			}
			if head.right != nil {
				ans.PushBack(strconv.Itoa(head.right.value))
				queue.PushBack(head.right)
			} else {
				ans.PushBack(nil)
			}
		}
	}
	return ans
}

func BuildByLevelQueue(levelList *list.List) *Node {
	if levelList == nil || levelList.Len() == 0 {
		return nil
	}
	head := generateNode(levelList.Front().Value)
	levelList.Remove(levelList.Front())

	queue := list.New()
	if head != nil {
		queue.PushBack(head)
	}
	var node *Node
	for queue.Len() != 0 {
		node = queue.Front().Value.(*Node)
		queue.Remove(queue.Front())

		node.left = generateNode(levelList.Front().Value)
		levelList.Remove(levelList.Front())
		node.right = generateNode(levelList.Front().Value)
		levelList.Remove(levelList.Front())
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
	return head
}

func generateNode(val any) *Node {
	v, ok := val.(string)
	if !ok {
		return nil
	}
	res, _ := strconv.Atoi(v)
	return NewNode(res)
}
