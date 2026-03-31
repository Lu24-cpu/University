package main

import (
	"fmt"
)

var f func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Operate(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	fmt.Println(Operate(3, 4, Add))
	fmt.Println(Operate(7, 10, Add))
}
