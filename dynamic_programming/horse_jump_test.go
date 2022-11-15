package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHorseJump(t *testing.T) {
	x := 7
	y := 7
	step := 10

	assert.Equal(t, HorseJump(x, y, step), HorseJumpDP(x, y, step))
}
