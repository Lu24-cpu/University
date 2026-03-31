package main

import (
	"fmt"
	"os"
	"strconv"
)

func rimuoviCifra(s string) []string {

	var slice []string

	slice = append(slice, s)

	var d int = 1

	for j := 0; j < 3; j++ {

		for n := range s {

			if n+d <= len(s) {

				st := s[:n] + s[n+d:]

				slice = append(slice, st)

			}

		}
		d++
	}

	return slice
}

func EPrimo(n int) bool {

	var d int = 1

	var count int

	for d <= n {

		if n%d == 0 {

			count++
		}
		d++
	}

	if count == 2 {

		return true
	}

	return false

}

func main() {

	slice := rimuoviCifra(os.Args[1])

	for _, i := range slice {

		num, _ := strconv.Atoi(i)

		if EPrimo(num) {

			fmt.Println(num)

		}
	}
	fmt.Println()

}
