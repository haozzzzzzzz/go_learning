package main

import (
	"fmt"
)

// 整型位操作类型
type BitUint32 uint32

// 判断某位是否为1
func (m BitUint32) IsSet1(index uint32) bool {
	return m&(1<<index) == (1 << index)
}

// 设置某位为1
func (m *BitUint32) SetBit1(index uint32) BitUint32 {
	*m = *m | (1 << index)
	return *m
}

// 设置某位为0
func (m *BitUint32) SetBit0(index uint32) BitUint32 {
	*m = *m & (0xFFFFFF ^ (1 << index))
	return *m
}

func main() {
	var bit BitUint32 = 2
	fmt.Println(bit.IsSet1(1))
	fmt.Println(bit.SetBit1(2))
	fmt.Println(bit.SetBit0(1))
}
