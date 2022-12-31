package string

func GetIndexOf(s1, s2 string) int {
	if len(s2) < 1 || len(s1) < len(s2) {
		return -1
	}
	str1 := []rune(s1)
	str2 := []rune(s2)
	x, y := 0, 0

	// O(M) m <= n，其中 m 表示 s2 字串長度，n 表示 s1 字串長度
	next := getNextArray(str2)

	// O(N)
	for x < len(str1) && y < len(str2) {
		if str1[x] == str2[y] {
			x++
			y++
		} else if next[y] == -1 { // y == 0
			x++
		} else {
			y = next[y]
		}
	}

	if y == len(str2) {
		return x - y
	} else {
		return -1
	}
}

func getNextArray(str2 []rune) []int {
	if len(str2) == 1 {
		return []int{-1}
	}
	next := make([]int, len(str2))
	next[0] = -1
	next[1] = 0
	i := 2  // 目前在哪個位置上求 next 切片的值
	cn := 0 // 目前是用哪個位置的值再和 i-1 位置的字元比較
	for i < len(next) {
		if str2[i-1] == str2[cn] { // 配成功的時候
			cn++
			next[i] = cn
			i++
		} else if next[cn] != -1 { // cn == 0
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
