package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var par, dis int

	if len(os.Args) == 1 {

		fmt.Printf("	dispari: 0, pari: 0\n")
		return
	}

	for _, i := range os.Args[1:] {

		numero, _ := strconv.Atoi(i)

		if numero%2 == 0 {

			par++
		} else {
			dis++
		}

	}

	fmt.Printf("	dispari: %d, pari: %d\n", dis, par)

}
