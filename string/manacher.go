package string

import "math"

func Manacher(s string) int {
	if len(s) == 0 {
		return 0
	}
	// "12132" -> "#1#2#1#3#2#"
	str := manacherString(s)
	// 回文半徑大小
	pArr := make([]int, len(str))
	// C 表示 R 更新過後的對稱點(回文中點)位置
	C := -1
	// R 表示最右的擴成功位置(最大回文半徑)的再下一個位置
	R := -1
	max := math.MinInt
	for i := 0; i < len(str); i++ {
		// R 第一個違規的位置，i >= R
		// i 位置擴出來的答案，i 位置擴的區域，至少是多大
		if R > i {
			// i 在 C 的對稱點位置: C - (i - C)
			pArr[i] = getMin(pArr[2*C-i], R-i)
		} else {
			pArr[i] = 1
		}
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		max = getMax(max, pArr[i])
	}
	return max - 1
}

func manacherString(str string) []rune {
	runeArr := []rune(str)
	res := make([]rune, len(str)*2+1)
	index := 0
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			res[i] = '#'
		} else {
			res[i] = runeArr[index]
			index++
		}
	}
	return res
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
