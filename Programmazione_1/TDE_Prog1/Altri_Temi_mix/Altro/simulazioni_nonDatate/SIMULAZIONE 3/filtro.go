package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {

	parole := os.Args[1:]

	var count int
	var flag bool

	for _, i := range parole {

		for _, lettera := range i {

			if unicode.IsDigit(lettera) {

				count++
				flag = true
				break

			}
		}

		if flag == false {

			fmt.Print(i, " ")
		} else {

			flag = false
		}

	}

	fmt.Println(count)

}
