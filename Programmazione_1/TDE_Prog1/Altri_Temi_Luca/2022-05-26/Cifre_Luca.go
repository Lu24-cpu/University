// Testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func contaMaxTot(s string) [10]int {
	var elementi [10]int
	var j int

	for _, elemento := range s {
		if unicode.IsNumber(elemento) && j <= 10 {
			n, _ := strconv.Atoi(string(elemento))
			elementi[n]++
			j++
		}
	}

	return elementi
}

func main() {
	riga := os.Args[1]
	fmt.Println(contaMaxTot(riga))
}

/*
Cifre
------

Scrivere un programma cifre.go dotato di
- una funzione
	contaCifre(s string) [10]int
che restituisce un array di lunghezza 10 che contiene il numero di
0, 1, 2, ... e 9 (in quest'ordine,) che si trovano nella stringa s.
Non ci sono vincoli sui tipi di caratteri in s.

- una funzione main() che legge una frase (tra "") da linea di comando che
contiene caratteri di qualsiasi tipo (cifre, lettere, simboli,
punteggiatura, lettere accentate, ecc.)
e stampa quante cifre di ogni tipo ci sono, ignorando gli eventuali altri caratteri.
Se manca l'argomento da linea di comando, deve stampare "manca argomento" e terminare.

Esempio
-------

Input: "1, 2, 3; non c’è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove."
Output: [0 2 2 2 1 1 1 1 1 1]
*/
