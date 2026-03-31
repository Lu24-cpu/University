package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	valori := os.Args[1:]

	var lettera, stringa string
	var numero int

	for indice, i := range valori {

		if indice%2 == 0 {

			lettera = i

		} else {

			numero, _ = strconv.Atoi(i)

			stringa += strings.Repeat(lettera, numero)
		}

	}

	fmt.Println(stringa)

}
