package binary_tree

import "fmt"

func unrecursivePreOrder(head *Node) {
	if head != nil {
		var stack = []*Node{head}
		for len(stack) != 0 {
			head = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", head.value)
			if head.right != nil {
				stack = append(stack, head.right)
			}
			if head.left != nil {
				stack = append(stack, head.left)
			}
		}
		fmt.Println()
	}
}

func unrecursiveInOrder(head *Node) {
	if head != nil {
		var stack []*Node
		for len(stack) != 0 || head != nil {
			if head != nil {
				stack = append(stack, head)
				head = head.left
			} else {
				head = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Printf("%d ", head.value)
				head = head.right
			}
		}
		fmt.Println()
	}
}

func unrecursivePostOrder1(head *Node) {
	if head != nil {
		var s1 = []*Node{head}
		var s2 []*Node
		for len(s1) != 0 {
			head = s1[len(s1)-1] // 從 s1 堆疊彈出的順序: 頭 右 左
			s1 = s1[:len(s1)-1]
			s2 = append(s2, head)
			if head.left != nil {
				s1 = append(s1, head.left)
			}
			if head.right != nil {
				s1 = append(s1, head.right)
			}
		}
		// 從 s2 堆疊彈出的順序: 左 右 頭
		for len(s2) != 0 {
			fmt.Printf("%d ", s2[len(s2)-1].value)
			s2 = s2[:len(s2)-1]
		}
		fmt.Println()
	}
}

func unrecursivePostOrder2(head *Node) {
	if head != nil {
		var stack = []*Node{head}
		var c *Node
		for len(stack) != 0 {
			c = stack[len(stack)-1]
			if c.left != nil && head != c.left && head != c.right {
				stack = append(stack, c.left)
			} else if c.right != nil && head != c.right {
				stack = append(stack, c.right)
			} else {
				fmt.Printf("%d ", c.value)
				stack = stack[:len(stack)-1]
				head = c
			}
		}
		fmt.Println()
	}
}
