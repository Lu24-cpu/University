package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	l, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < l; i++ {

		for j := 0; j < l; j++ {

			if i == j || i+j == l-1 {

				fmt.Print("*")
			} else {

				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}
