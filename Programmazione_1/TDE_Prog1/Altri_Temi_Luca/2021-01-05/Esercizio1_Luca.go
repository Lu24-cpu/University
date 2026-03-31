// Testo esercizio su pdf

package main

import (
	"fmt"
	"os"
	"unicode"
)

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	if posizione%2 == 0 {
		for i, carattere := range parola {
			if i%2 == 0 {
				parolaTrasformata += string(unicode.ToUpper(carattere))
			} else {
				parolaTrasformata += string(unicode.ToLower(carattere))
			}
		}
	} else {
		for i, carattere := range parola {
			if i%2 == 1 {
				parolaTrasformata += string(unicode.ToUpper(carattere))
			} else {
				parolaTrasformata += string(unicode.ToLower(carattere))
			}
		}
	}
	return
}

func main() {
	parole := os.Args[1:]

	for i, parola := range parole {
		fmt.Print(TrasformaParola(parola, i), " ")
	}
	fmt.Println()
}
