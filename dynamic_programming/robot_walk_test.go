package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobotWalk(t *testing.T) {
	ways1 := RobotWalk1(5, 2, 4, 6)
	ways2 := RobotWalk2(5, 2, 4, 6)
	ways3 := RobotWalk3(5, 2, 4, 6)

	assert.Equal(t, ways1, ways2)
	assert.Equal(t, ways2, ways3)
}
