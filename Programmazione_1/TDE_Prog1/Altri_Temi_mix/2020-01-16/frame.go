package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	dimensione, _ := strconv.Atoi(os.Args[1])
	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])

	var counti int
	var countj int = 1

	for i := 0; i < dimensione; i++ {

		for j := 0; j < dimensione; j++ {

			if i == 0 {

				fmt.Print(counti)

				counti++
			}

			if j == 0 && i > 0 {

				fmt.Print(countj)
				countj++
			}

			if i == x && j == y+1 {

				fmt.Print("*")
			} else {

				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
