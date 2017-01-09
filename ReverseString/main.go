// ReverseString project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseString("hello world"))
}

func reverseString(s string) string {
	var des []byte
	for i := len(s) - 1; i > -1; i-- {
		des = append(des, s[i])
	}
	return (string)(des)
}
