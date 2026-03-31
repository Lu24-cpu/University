package main

import (
	"fmt"
	"os"
)

func main() {

	parola := os.Args[1]

	mappa := make(map[string]int)

	slice := []string{}

	for i := 0; i < len(parola); i++ {

		for t := i + 1; t < len(parola); t++ {

			sottostringa := parola[i : t+1]

			if len(sottostringa) >= 3 && sottostringa[0] == sottostringa[len(sottostringa)-1] {

				mappa[sottostringa]++

				if mappa[sottostringa] == 1 {

					slice = append(slice, sottostringa)
				}

			}
		}

	}

	for i := 0; i < len(slice)-1; i++ {

		if len(slice[i]) < len(slice[i+1]) {

			slice[i], slice[i+1] = slice[i+1], slice[i]

			i = 0
		}

	}

	for _, i := range slice {

		fmt.Printf("%s -> Occorrenze: %d\n", i, mappa[i])
	}

}
