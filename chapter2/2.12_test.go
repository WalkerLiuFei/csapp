package chapter2

import (
	"testing"
	"fmt"
)

func Test212(t *testing.T){
	x := 0x87654321
	fmt.Printf("result a is 0x%x\n", x & 0xff)
	fmt.Printf("result b is 0x%x\n", x ^ 0xffffff00)
	fmt.Printf("result b is 0x%x\n", x | 0xff)
}

