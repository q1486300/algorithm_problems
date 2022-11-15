package binary_tree

import "fmt"

func PrintAllFolds(N int) {
	processPrintAllFolds(1, N, true)
	fmt.Println()
}

// 想像一棵二元樹，當你來到了一個節點
// 這個節點在第 i 層，一共有 N 層，N 固定不變
// 這個節點如果是凹的話，down = T
// 這個節點如果是凸的話，down = F
// 函數的功能: 中序印出想像的節點為根的整棵樹
func processPrintAllFolds(i, N int, down bool) {
	if i > N {
		return
	}
	processPrintAllFolds(i+1, N, true)
	if down {
		fmt.Print("凹 ")
	} else {
		fmt.Print("凸 ")
	}
	processPrintAllFolds(i+1, N, false)
}
