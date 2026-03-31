package main

import (
	"fmt"
)

func isBalanced(sequence string) bool {

	var c int

	for _, i := range sequence {

		if i == '(' {

			c++
		} else {

			c--

			if c < 0 {

				return false
			}
		}

	}

	if c == 0 {

		return true
	}

	return false

}

func sottostringhe(sequence string) (map[string]int, []string) {

	mappa := make(map[string]int)

	slice := []string{}

	for i := 0; i < len(sequence); i++ {

		for j := i; j < len(sequence); j++ {

			if len(sequence[i:j+1]) > 1 {

				if isBalanced(string(sequence[i : j+1])) {

					mappa[string(sequence[i:j+1])]++

					if mappa[string(sequence[i:j+1])] == 1 {

						slice = append(slice, sequence[i:j+1])
					}

				}
			}

		}
	}

	for i := 0; i < len(slice)-1; i++ {

		if len(slice[i]) > len(slice[i+1]) {

			slice[i], slice[i+1] = slice[i+1], slice[i]

			i = -1
		}
	}

	return mappa, slice
}

func main() {

	var s string

	fmt.Scan(&s)

	for _, i := range s {

		if i != '(' && i != ')' {

			return
		}
	}

	if isBalanced(s) {

		fmt.Println("è bilanciata")
	} else {

		fmt.Println("Non è bilanciata")
	}

	mappa, slice := (sottostringhe(s))

	for _, i := range slice {

		fmt.Println(i, mappa[i])
	}

}
