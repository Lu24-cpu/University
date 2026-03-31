package main

import (
	"fmt"
)

type ciao struct {
	x int
	y string
	z float64
}

func (x ciao) Stampa() string {
	return fmt.Sprintf("%s, %.3f, %d", x.y, x.z, x.x)
}

func main() {
	var a ciao

	fmt.Println(a.Stampa())
}
