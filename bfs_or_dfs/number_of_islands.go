package bfs_or_dfs

import (
	"container/list"
)

// 測試連結: https://leetcode.com/problems/number-of-islands/
func NumIslands2(board [][]byte) int {
	islands := 0
	for i, ints := range board {
		for j, value := range ints {
			if value == 1 {
				islands++
				infect(board, i, j)
			}
		}
	}
	return islands
}

// 從 (i, j) 位置出發，把所有連成一片的 1 變成 0
func infect(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != 1 {
		return
	}
	board[i][j] = 0
	infect(board, i-1, j)
	infect(board, i+1, j)
	infect(board, i, j-1)
	infect(board, i, j+1)
}

func NumIslands1(board [][]byte) int {
	row := len(board)
	col := len(board[0])

	dots := make([][]*Dot, row)
	for i := range dots {
		dots[i] = make([]*Dot, col)
	}
	var dotList []*Dot

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 1 {
				dots[i][j] = NewDot(i, j)
				dotList = append(dotList, dots[i][j])
			}
		}
	}

	uf := NewUnionFind(dotList)
	for j := 1; j < col; j++ {
		if board[0][j-1] == 1 && board[0][j] == 1 {
			uf.Union(dots[0][j-1], dots[0][j])
		}
	}
	for i := 1; i < row; i++ {
		if board[i-1][0] == 1 && board[i][0] == 1 {
			uf.Union(dots[i-1][0], dots[i][0])
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == 1 {
				if board[i][j-1] == 1 {
					uf.Union(dots[i][j-1], dots[i][j])
				}
				if board[i-1][j] == 1 {
					uf.Union(dots[i-1][j], dots[i][j])
				}
			}
		}
	}
	return uf.Sets()
}

type Dot struct {
	row, col int
}

func NewDot(row, col int) *Dot {
	return &Dot{
		row: row,
		col: col,
	}
}

type UnionFind[V comparable] struct {
	parents map[V]V
	sizeMap map[V]int
}

func NewUnionFind[V comparable](values []V) UnionFind[V] {
	uf := UnionFind[V]{
		parents: make(map[V]V),
		sizeMap: make(map[V]int),
	}
	for _, cur := range values {
		uf.parents[cur] = cur
		uf.sizeMap[cur] = 1
	}
	return uf
}

func (uf UnionFind[V]) FindFather(cur V) V {
	path := list.New()
	for cur != uf.parents[cur] {
		path.PushBack(cur)
		cur = uf.parents[cur]
	}
	for path.Len() != 0 {
		uf.parents[path.Back().Value.(V)] = cur
		path.Remove(path.Back())
	}
	return cur
}

func (uf UnionFind[V]) Union(a, b V) {
	aHead := uf.FindFather(a)
	bHead := uf.FindFather(b)
	if aHead != bHead {
		aSetSize := uf.sizeMap[aHead]
		bSetSize := uf.sizeMap[bHead]
		if aSetSize >= bSetSize {
			uf.parents[bHead] = aHead
			uf.sizeMap[aHead] = aSetSize + bSetSize
			delete(uf.sizeMap, bHead)
		} else {
			uf.parents[aHead] = bHead
			uf.sizeMap[bHead] = aSetSize + bSetSize
			delete(uf.sizeMap, aHead)
		}
	}
}

func (uf UnionFind[V]) Sets() int {
	return len(uf.sizeMap)
}
