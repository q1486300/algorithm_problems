package bitmap

import (
	"bytes"
	"fmt"
)

type BitMap []byte

const byteSize = 8

func NewBitMap(maxNum uint) BitMap {
	return make(BitMap, (maxNum/byteSize)+1)
}

func (b BitMap) Set(n uint) {
	if n/byteSize >= uint(len(b)) {
		fmt.Println("大小超出 bitmap 範圍")
		return
	}
	byteIndex := n / byteSize   // 第 x 個 Byte (0, 1, 2...)
	offsetIndex := n % byteSize // 偏移量 (0 <= 偏移量 < byteSize)
	// 第 x 個 Byte 偏移量為 offsetIndex 的位，或(|)上 1
	b[byteIndex] |= 1 << offsetIndex
}

func (b BitMap) Del(n uint) {
	if n/byteSize >= uint(len(b)) {
		fmt.Println("大小超出 bitmap 範圍")
		return
	}
	byteIndex := n / byteSize
	offsetIndex := n % byteSize
	// 第 x 個 Byte 偏移量為 offsetIndex 的位，與(&)上 0
	b[byteIndex] &= ^(1 << offsetIndex)
}

func (b BitMap) IsExist(n uint) bool {
	if n/byteSize >= uint(len(b)) {
		fmt.Println("大小超出 bitmap 範圍")
		return false
	}
	byteIndex := n / byteSize
	offsetIndex := n % byteSize
	return b[byteIndex]&(1<<offsetIndex) != 0
}

func (b BitMap) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("BitMap {")
	for i, v := range b {
		if v == 0 {
			continue
		}
		for j := 0; j < byteSize; j++ {
			if v&(1<<j) != 0 {
				fmt.Fprintf(&buf, "%d ", byteSize*i+j)
			}
		}
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString("}\n")
	return buf.String()
}
