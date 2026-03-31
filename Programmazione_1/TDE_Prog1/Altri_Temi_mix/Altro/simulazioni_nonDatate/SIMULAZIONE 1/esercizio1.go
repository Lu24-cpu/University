package main

import (
	"fmt"
	"os"
	"unicode"
)

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {

	var c int

	for indice, i := range parola {

		if posizione%2 == 0 {

			if unicode.IsUpper(i) == false && indice%2 == 0 {

				parolaTrasformata += string(i - 32)
				c++

			}

		} else {

			if unicode.IsUpper(i) == false && indice%2 != 0 {

				parolaTrasformata += string(i - 32)
				c++

			}

		}

		if c == 0 {

			parolaTrasformata += string(i)
		} else {

			c--
		}
	}

	return parolaTrasformata
}

func main() {
	parole := os.Args[1:]

	fmt.Println("ciao")

	for pos, i := range parole {

		fmt.Print(TrasformaParola(i, pos) + " ")

	}
	fmt.Println()

}
