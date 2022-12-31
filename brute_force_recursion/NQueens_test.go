package brute_force_recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNQueens(t *testing.T) {
	n := 14

	start := time.Now()
	n2 := NQueens2(n)
	duration := time.Since(start)
	fmt.Printf("NQueens2 cost time: %v\n", duration)

	start = time.Now()
	n1 := NQueens1(n)
	duration = time.Since(start)
	fmt.Printf("NQueens1 cost time: %v\n", duration)

	assert.Equal(t, n1, n2)
}
