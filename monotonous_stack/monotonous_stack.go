package monotonous_stack

import (
	"container/list"
)

// arr = [ 3, 1, 2, 3]
//		   0  1  2  3
// [
//		0: [-1,  1]
//		1: [-1, -1]
//		2: [ 1, -1]
//		3: [ 2, -1]
// ]
func GetNearLessNoRepeat(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	// 只存索引
	stack := list.New()
	for i, value := range arr { // 當遍歷到 i 位置的數，arr[i]
		for stack.Len() != 0 && arr[stack.Back().Value.(int)] > value {
			j := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			leftLessIndex := -1
			if stack.Len() != 0 {
				leftLessIndex = stack.Back().Value.(int)
			}
			res[j][0] = leftLessIndex
			res[j][1] = i
		}
		stack.PushBack(i)
	}
	for stack.Len() != 0 {
		j := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		leftLessIndex := -1
		if stack.Len() != 0 {
			leftLessIndex = stack.Back().Value.(int)
		}
		res[j][0] = leftLessIndex
		res[j][1] = -1
	}
	return res
}

func GetNearLess(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	stack := list.New()

	for i, value := range arr { // i -> arr[i] 進堆疊
		if stack.Len() == 0 {
			newList := list.New()
			newList.PushBack(i)
			stack.PushBack(newList)
			continue
		}
		backList := stack.Back().Value.(*list.List)
		lastIndex := backList.Back().Value.(int)
		for stack.Len() != 0 && arr[lastIndex] > value {
			stack.Remove(stack.Back())
			leftLessIndex := -1
			var backbackList *list.List
			if stack.Len() != 0 {
				backbackList = stack.Back().Value.(*list.List)
				leftLessIndex = backbackList.Back().Value.(int)
				lastIndex = leftLessIndex
			}
			for backList.Len() != 0 {
				index := backList.Front().Value.(int)
				backList.Remove(backList.Front())
				res[index][0] = leftLessIndex
				res[index][1] = i
			}
			backList = backbackList
		}
		if arr[lastIndex] == value {
			backList.PushBack(i)
		} else {
			newList := list.New()
			newList.PushBack(i)
			stack.PushBack(newList)
		}
	}

	for stack.Len() != 0 {
		backList := stack.Back().Value.(*list.List)
		stack.Remove(stack.Back())
		leftLessIndex := -1
		if stack.Len() != 0 {
			backbackList := stack.Back().Value.(*list.List)
			leftLessIndex = backbackList.Back().Value.(int)
		}
		for backList.Len() != 0 {
			index := backList.Front().Value.(int)
			backList.Remove(backList.Front())
			res[index][0] = leftLessIndex
			res[index][1] = -1
		}
	}
	return res
}
