package bfs_or_dfs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestNumIslands(t *testing.T) {
	row := 1000
	col := 1000
	board1 := generateRandomMatrix(row, col)
	board2 := copyMatrix(board1)

	// 感染方法，並查集(map 實現)
	start := time.Now()
	n1 := NumIslands1(board1)
	duration := time.Since(start)
	fmt.Printf("感染方法，並查集(map 實現)的執行時間: %v\n", duration)

	// 感染方法，普通 BFS
	start = time.Now()
	n2 := NumIslands2(board2)
	duration = time.Since(start)
	fmt.Printf("感染方法，普通 BFS 的執行時間: %v\n", duration)

	fmt.Println(n1, n2)
	assert.Equal(t, n1, n2)
}

func generateRandomMatrix(row, col int) [][]byte {
	board := make([][]byte, row)
	for i := range board {
		board[i] = make([]byte, col)
	}

	for i, bytes := range board {
		for j := range bytes {
			if rand.Float32() < 0.5 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
	return board
}

func copyMatrix(board [][]byte) [][]byte {
	result := make([][]byte, len(board))
	for i := range result {
		result[i] = make([]byte, len(board[i]))
		copy(result[i], board[i])
	}
	return result
}
