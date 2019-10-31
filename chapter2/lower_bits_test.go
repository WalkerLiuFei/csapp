package chapter2

import (
	"testing"
	"fmt"
)

func TestLowBits(t *testing.T) {
	n := int64(17)
	fmt.Printf("result of  n := %d result is %x", n, lowBits2(n))
}

func lowBits2(n int64) int {
	max := uint(n - 1 - n)
	fmt.Printf("result of  n := %x", max)
	return int(max >> uint(64-n))
}
func lowBits(n int) int {
	result := 0
	for n = n - 1; n >= 0; n-- {
		result = result | (0x01 << uint(n))
	}
	return result
}
