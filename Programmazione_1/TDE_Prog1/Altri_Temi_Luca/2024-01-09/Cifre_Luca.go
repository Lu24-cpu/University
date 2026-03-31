// Testo a fine Esercizio
package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func contaMaxTot(s string) ([10]int, int, int) {
	var elementi [10]int
	var max, sum, j int

	for _, elemento := range s {
		if unicode.IsNumber(elemento) && j <= 10 {
			n, _ := strconv.Atoi(string(elemento))
			elementi[n]++
			if max < n {
				max = n
			}
			sum += n
		}
	}

	return elementi, max, sum
}

func main() {
	riga := os.Args[1]
	fmt.Println(contaMaxTot(riga))
}

/*
Cifre
------

# Scrivere un programma cifre.go che elabora le cifre contenute in un testo passato da linea di comando.

Il programma deve essere dotato di:

  - una funzione
    contaMaxTot(s string) ([10]int, int, int)

che restituisce, in quest'ordine:
  - un array di lunghezza 10 che contiene il numero di '0', '1', '2', ... e '9' (in quest'ordine,) che si trovano nella stringa s,
  - la massima cifra presente nella stringa s (se non sono presenti cifre nella stringa, il valore restituito deve essere -1)
  - la somma delle cifre contenute nella stringa s.

Non ci sono vincoli sui tipi di caratteri in s (possono cioè essere anche non ASCII).

- una funzione main() che legge una frase (racchiusa tra "") da linea di comando che
contiene caratteri di qualsiasi tipo (cifre, lettere, simboli, punteggiatura, lettere accentate, ecc.)
e stampa quante cifre di ogni tipo ci sono, qual è la cifra più grande e la somma delle cifre incontrate.
Se manca l'argomento da linea di comando, deve stampare "manca argomento" e terminare.

Esempio
-------

Input: "1, 2, 3; non c'è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove."
Output: [0 2 2 2 1 1 1 1 1 1] 9 51
*/
