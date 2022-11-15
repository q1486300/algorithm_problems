package bit

func GetMax(a, b int) int {
	c := a - b
	sa := sign(a)
	sb := sign(b)
	sc := sign(c)
	difSab := sa ^ sb                 // a 和 b 的符號不一樣為 1，一樣為 0
	sameSab := flip(difSab)           // a 和 b 的符號一樣為 1，不一樣為 0
	returnA := difSab*sa + sameSab*sc // a 和 b 符號不一樣且 a 非負數，或者a 和 b 符號一樣且 c 非負數(表示 a 比較大)，returnA 為 1
	returnB := flip(returnA)          // returnA 的互斥條件
	return a*returnA + b*returnB      // 互斥條件的加法取代 if-else 條件判斷
}

// n 非負數，返回 1
// n 負數，返回 0
func sign(n int) int {
	return flip((n >> 31) & 1)
}

// 1 -> 0
// 0 -> 1
func flip(n int) int {
	return n ^ 1
}
