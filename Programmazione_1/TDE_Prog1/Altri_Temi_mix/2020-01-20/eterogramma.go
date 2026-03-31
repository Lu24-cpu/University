package main

import (
	"fmt"
	"os"
)

func main() {

	mappa := make(map[rune]int)

	for _, i := range os.Args[1] {

		if i != ' ' {

			mappa[i]++

			if mappa[i] > 1 {

				fmt.Println("Non è un eterogramma")
				return
			}
		}

	}

	fmt.Println("È un eterogramma")

}
