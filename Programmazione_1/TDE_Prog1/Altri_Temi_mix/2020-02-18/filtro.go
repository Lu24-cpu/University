package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	input := os.Args[1:]
	var count int

	for _, i := range input {

		numero, _ := strconv.Atoi(i)
		count = 0

		for {

			if count > 3 {

				break
			}

			if math.Pow(3, float64(count)) == float64(numero) {

				fmt.Print(numero, " ")

				break
			} else {

				count++
			}

		}

	}
	fmt.Println()

}
