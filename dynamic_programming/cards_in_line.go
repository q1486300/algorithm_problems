package dynamic_programming

import "math"

// 此題的暴力遞迴版本在: brute_force_recursion/cards_in_line.go
func Win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	N := len(arr)
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := range fmap {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		for j := range fmap[i] {
			fmap[i][j] = -1
			gmap[i][j] = -1
		}
	}
	first := f1(arr, 0, N-1, fmap, gmap)
	second := g1(arr, 0, N-1, fmap, gmap)
	return int(math.Max(float64(first), float64(second)))
}

// arr[L..R]，返回先手獲得的最好分數
func f1(arr []int, L, R int, fmap, gmap [][]int) int {
	if fmap[L][R] != -1 {
		return fmap[L][R]
	}
	ans := 0
	if L == R {
		ans = arr[L]
	} else {
		p1 := arr[L] + g1(arr, L+1, R, fmap, gmap)
		p2 := arr[R] + g1(arr, L, R-1, fmap, gmap)
		ans = int(math.Max(float64(p1), float64(p2)))
	}
	fmap[L][R] = ans
	return ans
}

// arr[L..R]，返回後手獲得的最好分數
func g1(arr []int, L, R int, fmap, gmap [][]int) int {
	if gmap[L][R] != -1 {
		return gmap[L][R]
	}
	ans := 0
	if L != R {
		p1 := f1(arr, L+1, R, fmap, gmap)
		p2 := f1(arr, L, R-1, fmap, gmap)
		ans = int(math.Min(float64(p1), float64(p2)))
	}
	gmap[L][R] = ans
	return ans
}

func Win2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	N := len(arr)
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := range fmap {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
	}
	for i := range fmap {
		fmap[i][i] = arr[i]
	}
	for startCol := 1; startCol < N; startCol++ {
		L := 0
		R := startCol
		for R < N {
			fmap[L][R] = int(math.Max(float64(arr[L]+gmap[L+1][R]), float64(arr[R]+gmap[L][R-1])))
			gmap[L][R] = int(math.Min(float64(fmap[L+1][R]), float64(fmap[L][R-1])))
			L++
			R++
		}
	}
	return int(math.Max(float64(fmap[0][N-1]), float64(gmap[0][N-1])))
}
