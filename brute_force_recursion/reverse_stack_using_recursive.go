package brute_force_recursion

func Reverse(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	i := popBottom(stack)
	Reverse(stack)
	*stack = append(*stack, i)
}

// 堆疊底部元素移除掉
// 上面的元素蓋下來
// 返回移除掉的堆疊底部元素
func popBottom(stack *[]int) int {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	if len(*stack) == 0 {
		return result
	} else {
		last := popBottom(stack)
		*stack = append(*stack, result)
		return last
	}
}
