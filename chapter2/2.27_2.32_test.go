package chapter2

import (
	"testing"
	"fmt"
	"math"
)

func Test2_27(t *testing.T) {
	fmt.Println(uadd_ok(1111, 2))
	fmt.Println(tadd_ok(5555, 2))
	fmt.Println(tsub_ok(math.MaxInt32, math.MinInt32))
	fmt.Println(tmul_ok(math.MaxInt32, 10))
	fmt.Println(tmul_ok(10, 10))
}

func uadd_ok(a, b uint32) bool {

	return a+b > a && a+b > b
}

func tadd_ok(a, b int32) bool {
	return !(a+b < 0)
}

func tsub_ok(a, b int32) bool {
	if b < 0 {
		return a-b > 0
	}
	if b > 0 {
		return (a - b) < a
	}
	return true
}

func tmul_ok(a, b int32) bool {
	return int64(a*b)>>32 == 0
}
