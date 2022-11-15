package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialAndReconstruct(t *testing.T) {
	testTimes := 1000000
	maxLevel := 5
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBST(maxLevel, maxValue)

		pre := PreSerial(head)
		post := PostSerial(head)
		level := LevelSerial(head)
		preBuild := BuildByPreQueue(pre)
		postBuild := BuildByPostQueue(post)
		levelBuild := BuildByLevelQueue(level)

		if !assert.Equal(t, preBuild, postBuild) || !assert.Equal(t, postBuild, levelBuild) {
			PrintTree(preBuild)
			PrintTree(postBuild)
			PrintTree(levelBuild)
			break
		}
	}
}
