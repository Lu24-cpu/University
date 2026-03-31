// Testo esercizio su pdf

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	stringa := os.Args[1:]

	for i := 1; i < len(stringa); i++ {
		n, _ := strconv.Atoi(stringa[i])

		for j := 0; j < n; j++ {
			fmt.Print(stringa[i-1])
		}
		i++
	}
	fmt.Println()
}
