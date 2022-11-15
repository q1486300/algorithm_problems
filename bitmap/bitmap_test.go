package bitmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap := NewBitMap(20)
	bitMap.Set(20)
	bitMap.Set(1)
	bitMap.Set(2)
	bitMap.Set(3)
	bitMap.Set(4)
	bitMap.Set(5)

	assert.Equal(t, true, bitMap.IsExist(20))
	assert.Equal(t, true, bitMap.IsExist(5))

	bitMap.Del(20)

	assert.Equal(t, false, bitMap.IsExist(20))
	assert.Equal(t, true, bitMap.IsExist(1))

	bitMap.Del(1)

	assert.Equal(t, false, bitMap.IsExist(1))
	assert.Equal(t, true, bitMap.IsExist(2))

	bitMap.Del(3)

	assert.Equal(t, false, bitMap.IsExist(3))
	assert.Equal(t, true, bitMap.IsExist(4))
	assert.Equal(t, true, bitMap.IsExist(5))

	fmt.Println(bitMap)
}
