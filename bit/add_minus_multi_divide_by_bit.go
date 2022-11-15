package bit

func Add(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b      // 無進位相加的結果
		b = (a & b) << 1 // 進位信息
		a = sum
	}
	return sum
}

func Minus(a, b int) int {
	return Add(a, negNum(b))
}

func negNum(n int) int {
	return Add(^n, 1)
}

func Multi(a, b int) int {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res = Add(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

func Divide(a, b int) int {
	x := a
	if isNeg(a) {
		x = negNum(a)
	}
	y := b
	if isNeg(b) {
		y = negNum(b)
	}

	res := 0
	for i := 31; i > -1; i = Minus(i, 1) {
		if (x >> i) >= y {
			res |= 1 << i
			x = Minus(x, y<<i)
		}
	}
	if isNeg(a) && !isNeg(b) || !isNeg(a) && isNeg(b) {
		return negNum(res)
	} else {
		return res
	}
}

func isNeg(n int) bool {
	return n < 0
}
