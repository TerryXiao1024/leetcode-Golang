// HammingDistance project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(4, 4))
}

func hammingDistance(x int, y int) int {
	var res int
	var i uintptr
	exc := x ^ y
	for i = 0; i < 32; i++ {
		res = res + (exc>>i)&1
	}
	return res
}
