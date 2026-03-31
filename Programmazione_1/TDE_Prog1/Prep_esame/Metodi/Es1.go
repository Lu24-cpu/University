// Crea un codice che permetta la stampa di uno array/slice senza le parentesi quadre a inizio e fine

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stringa string

const DISTANZA = 33

func (riga Stringa) Maiuscolo() (uppercase string) {
	for _, el := range riga {
		uppercase += string(el - DISTANZA)
	}
	return
}

func (riga Stringa) Minuscolo() (lowercase string) {
	for _, el := range riga {
		lowercase += string(el + DISTANZA)
	}
	return
}

func main() {
	var testo []Stringa
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo = append(testo, Stringa(scanner.Text()))
	}

	for _, riga := range testo {
		fmt.Println(riga.Maiuscolo())
	}
	for _, riga := range testo {
		fmt.Println(riga.Minuscolo())
	}
}
