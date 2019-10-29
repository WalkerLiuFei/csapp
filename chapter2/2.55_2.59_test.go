package chapter2

import (
	"fmt"
	"unsafe"
	"testing"
	"math"
)

//---------------------2.57----------------

func showBytes(n int, size int) {
	for count := size -1; count >= 0; count-- {
		fmt.Printf("%.2X ",byte(n >> uint(8 * count)))
	}
}

type bit32Data [4]byte
type bit64Data [4]byte
// this is not safe , the size of n must be small than 32 bits,otherwise it will panic!
func showBytes2(n int, size int) {
	nBytes :=  *(*bit32Data)(unsafe.Pointer(&n))
	fmt.Println(nBytes)
	for index := size -1; index >= 0; index-- {
		fmt.Printf("%.2X ",nBytes[index])
	}
}

func TestShowInt16(t *testing.T) {
	number := int16(math.MaxInt16)
	//showBytes(int(number),(int)(unsafe.Sizeof(number)))
	showBytes2(int(number),(int)(unsafe.Sizeof(number)))
}

func TestShowFloat32(t *testing.T) {
	number := float32(math.MaxFloat32)
	showBytes(int(math.Float32bits(number)),(int)(unsafe.Sizeof(number)))
}

func TestShowFloat64(t *testing.T) {
	number := float64(math.MaxFloat64)
	showBytes(int(math.Float64bits(number)),(int)(unsafe.Sizeof(number)))
}



//---------------------2.58----------------
func TestIsLittleEndian(t *testing.T){
	testNum := int16(0xff)
	toBytes :=  *(*bit32Data)(unsafe.Pointer(&testNum))
	if toBytes[0] == 0xff{
		//true
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
	//false
}

//---------------------2.59----------------
func TestMergeNumer(t *testing.T){
	x := 0x89ABCDEF
	y := 0x76543210
	fmt.Printf("Result is %X",(x & 0x000000ff) | (y & 0xffffff00))
}

