package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("输入一个整数")
	fmt.Scanf("%d", &a)
	i := 2
	for i = 2; i <= a-1; i++ {
		if a%i == 0 {
			break
		}
	}
	if i >= a {
		println("a是素数")
	} else {
		println("a不是素数")
	}
}
