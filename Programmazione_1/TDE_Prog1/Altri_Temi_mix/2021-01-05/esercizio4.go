package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := os.Args[1]

	d, _ := strconv.Atoi(os.Args[2])

	

	slice := []string{}
	slice1 := []int{}

	var s string

	for _, i := range n {

		slice = append(slice, string(i))
	}

	

	fmt.Println(slice)

}
