package chapter2

import (
	"testing"
	"fmt"
)

func TestRotateRight(t *testing.T) {
	a := 0x12345678
	n := 20
	fmt.Printf("rotate  %x by %d is %x",a,n,rotateRight(uint32(a),uint32(n)))
}

func rotateRight(x uint32, n uint32) uint32 {
	// 取出低n 位
	x1 := (1<<32-1)>> (32-n) & x
	// 将低n位挪到前面 ，将高位的移到后面
	return (x1 << (32 - n)) | (x >> n)
}
