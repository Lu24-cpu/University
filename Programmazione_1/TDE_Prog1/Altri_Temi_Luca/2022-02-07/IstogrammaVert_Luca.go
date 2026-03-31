// testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
)

func Maggiore(q []int) (max int) {
	for _, el := range q {
		if el > max {
			max = el
		}
	}
	return
}

func main() {
	var qint []int
	q := os.Args[1:]

	for _, elemento := range q {
		n, _ := strconv.Atoi(elemento)
		qint = append(qint, n)
	}
	mag := Maggiore(qint)

	for i := 0; i < mag; i++ {
		for _, el := range qint {
			if mag-el <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/*
ISTOGRAMMA
----------

Scrivere un programma (il file deve chiamarsi 'istogrammaVert.go') che legge da linea di comando una serie di numeri interi
non negativi e ne disegna il relativo istogramma. L'istogramma sarà formato da colonne di asterischi che rappresentano i valori letti.
Non sono richiesti controlli sui dati.

Esempio:

$ ./istogrammaVert 5 0 2 9 4
   *
   *
   *
   *
*  *
*  **
*  **
* ***
* ***

$ ./istogrammaVert 2 3 4 5
   *
  **
 ***
****
****

*/
