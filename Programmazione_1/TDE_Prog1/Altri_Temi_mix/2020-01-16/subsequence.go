package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var somma int

	var count int = 1

	slice := []int{}

	for i := 1; i < len(os.Args); i++ {

		numero, _ := strconv.Atoi(os.Args[i])

		somma += numero

		if somma != 0 {

			slice = append(slice, numero)

		} else {

			slice = append(slice, numero)

			fmt.Println(slice)

			slice = nil

			somma = 0

			count++

			i = count
		}

	}

}
