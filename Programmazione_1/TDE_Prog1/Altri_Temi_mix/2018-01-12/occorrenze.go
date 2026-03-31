package main

import (
	"fmt"
	"os"
)

func main() {
	n := os.Args[1]
	r := n[0]
	s := os.Args[2]
	occ := occorrenze(s, r, 0, 0)
	fmt.Println("Occorrenze di", n, "in", s, "\n", occ)
}

func occorrenze(s string, r byte, i int, count int) int {
	if i == len(s) {
		return count
	}
	if s[i] == r {
		count++
	}
	i++
	return occorrenze(s, r, i, count)
}
