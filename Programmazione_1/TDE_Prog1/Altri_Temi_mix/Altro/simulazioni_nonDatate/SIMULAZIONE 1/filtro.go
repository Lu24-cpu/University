package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var lettera string
	var numero int

	for i := 1; i < len(os.Args); i++ {

		if i%2 != 0 {

			lettera = os.Args[i]

		} else {

			numero, _ = strconv.Atoi(os.Args[i])
			fmt.Print(strings.Repeat(lettera, numero))
		}

	}
	fmt.Println()
}
