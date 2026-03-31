package main

import (
	"fmt"
	"os"
	"unicode"
)

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {

	if posizione%2 == 0 {

		for indice, i := range parola {

			if indice%2 == 0 {

				if unicode.IsUpper(i) == false {

					parolaTrasformata += string(i - 32)
				}
			} else {

				parolaTrasformata += string(i)
			}

		}

		return parolaTrasformata
	}

	for indice, i := range parola {

		if indice%2 != 0 {

			if unicode.IsUpper(i) == false {

				parolaTrasformata += string(i - 32)
			}
		} else {

			parolaTrasformata += string(i)
		}

	}

	return parolaTrasformata

}
func main() {

	parole := os.Args[1:]

	fmt.Println(parole)

	for indice, i := range parole {

		fmt.Print(TrasformaParola(string(i), indice) + " ")
	}

}
