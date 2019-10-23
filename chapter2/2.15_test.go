package chapter2

import (
	"testing"
	"fmt"
)

/**

 */

func TestEqual(t *testing.T){
	fmt.Printf("Result is :%t",equal(10,10))
}
func equal(x,y int) bool {
	// right way in c : !(x ^ y)
	return x ^ y == 0
}