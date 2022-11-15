package sliding_window

import "container/list"

func GetMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	// qMax 視窗最大值的更新結構
	// 放索引
	qMax := list.New()
	res := make([]int, len(arr)-w+1)
	index := 0
	for R := 0; R < len(arr); R++ {
		for qMax.Len() != 0 && arr[qMax.Back().Value.(int)] <= arr[R] {
			qMax.Remove(qMax.Back())
		}
		qMax.PushBack(R)
		if qMax.Front().Value.(int) == R-w { // R - w 過期的索引
			qMax.Remove(qMax.Front())
		}
		if R >= w-1 { // 視窗形成了
			res[index] = arr[qMax.Front().Value.(int)]
			index++
		}
	}
	return res
}
