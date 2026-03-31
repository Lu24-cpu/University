package main

import (
	"fmt"
)

type Matrice [][]int

func main() {
	var a, b Matrice

	a = Matrice{[]int{0, 1, 2, 3}, []int{4, 5, 6, 7}, []int{8, 9, 10, 11}}
	b = Matrice{[]int{12, 13, 14, 15}, []int{16, 17, 18, 19}, []int{20, 21, 22, 23}}

	c := Somma(a, b)

	for i := range c {
		for _, n := range c[i] {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
}

func Somma(a, b Matrice) Matrice {
	if len(a) == 0 && len(b) == 0 {
		return Matrice{}
	}
	x := len(a) - 1
	return append(Somma(a[:x], b[:x]), SommaSl(a[x], b[x]))
}

func SommaSl(a, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}
	x := len(a) - 1
	return append(SommaSl(a[:x], b[:x]), a[x]+b[x])
}
