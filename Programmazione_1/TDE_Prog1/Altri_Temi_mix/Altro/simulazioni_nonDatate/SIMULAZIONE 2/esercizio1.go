package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	parole := os.Args[1:]

	var somma, num int

	var err error

	for indice, i := range parole {

		if indice%2 == 0 {

			num, err = strconv.Atoi(i)

			if err == nil {

				if num <= somma {

					fmt.Println("Valore in posizione", indice, "non valido")
					return
				}

				somma += num
			}

		}
	}

	fmt.Println("Sequenza valida")
}
