package brute_force_recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNQueens(t *testing.T) {
	n := 14

	start := time.Now().UnixMilli()
	n2 := NQueens2(n)
	end := time.Now().UnixMilli()
	fmt.Printf("NQueens2 cost time: %d ms\n", end-start)

	start = time.Now().UnixMilli()
	n1 := NQueens1(n)
	end = time.Now().UnixMilli()
	fmt.Printf("NQueens1 cost time: %d ms\n", end-start)

	assert.Equal(t, n1, n2)
}
