// testo esercizio a fine Codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		frase := scanner.Text()
		parole := strings.Fields(frase)

		for _, parola := range parole {
			caratteri := []rune(parola)
			if parola == "stop" {
				return
			}

			if len(caratteri)%2 == 0 && parola != "stop" {
				for i := len(parola) - 1; i >= 0; i-- {
					fmt.Printf("%c", caratteri[i])
				}
				fmt.Println()
			}
		}
	}
	fmt.Println()
}

/*

Revert
------

Realizzare un programma Go (nome file 'revert.go') che legge da standard input **stringhe** costituite da soli caratteri
**ASCII standard (a 7 bit)** (separate da un numero arbitrario di white space) e stampa solo quelle di lunghezza pari, al contrario
(il primo carattere per ultimo e l'ultimo per primo, ecc.), una per riga.
Il programma termina quando legge la stringa "stop", che non va trattata.

*/
