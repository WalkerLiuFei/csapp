package chapter2

import (
	"testing"
	"fmt"
	"unsafe"
)

func TestReplaceByte(t *testing.T) {
	target := int32(0xABCDEF)
	result := replace_byte2(int(target), 0x12, 1)
	fmt.Printf("Result is %X", result)
}

// use unsafe pointer converter must have special size type.
func replace_byte(x int32, b byte, i int) int32 {
	bytesData := *(*bit32Data)(unsafe.Pointer(&x))
	bytesData[i] = bytesData[i]&0x00 | b
	return *(*int32)(unsafe.Pointer(&bytesData))
}

func replace_byte2(x int, b byte, i int) int {
	temp := 0xFF
	temp = temp << uint(8*i)

	return (x & ^temp) | int(b)<<uint(8*1)
}

func Test2_61(t *testing.T) {
	t2_61(11)
}
func t2_61(x uint32) uint32 {

	fmt.Printf("%X", 0xffffffff)
	/*^x ^ 0xffffffff
	return x | 0xff000000 */
	// 所有位等1 ,取反得 0
	a := ^x
	// 所有位等0 取反得 ~(0xffffffff & 0xffffffff),得0
	b := ^(^x & 0xffffffff)
	//最低字节为0，那么结果为0
	c := x & 0x000000ff
	// 最高字节都为1，结果为0
	d := ^(x & 0xff000000) & 0xff000000
	return a | b | c | d
}
