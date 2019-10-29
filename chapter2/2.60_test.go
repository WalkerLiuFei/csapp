package chapter2

import (
	"testing"
	"fmt"
	"unsafe"
)

func TestReplaceByte(t *testing.T) {
	target :=  int32(0xABCDEF)
	result := replace_byte2(int(target), 0x12, 1)
	fmt.Printf("Result is %X",result)
}
// use unsafe pointer converter must have special size type.
func replace_byte(x int32, b byte, i int) int32 {
	bytesData := *(*bit32Data)(unsafe.Pointer(&x))
	bytesData[i] = bytesData[i] & 0x00 | b
	return *(*int32)(unsafe.Pointer(&bytesData))
}

func replace_byte2(x int, b byte, i int) int {
	temp := 0xFF
	temp = temp << uint(8 * i)

	return (x & ^temp) | int(b) << uint(8 * 1)
}
