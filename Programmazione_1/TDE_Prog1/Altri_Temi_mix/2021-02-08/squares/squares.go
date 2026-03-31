package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NestedSquares(n int) string {

	if n < 5 {

		return ""
	}

	riga := strings.Repeat("*", n)                       //prima e ultima riga del quadrato fuori
	riga2 := "*" + strings.Repeat(" ", n-2) + "*" + "\n" //seconda e penultima riga del quadrato fuori
	riga3 := "* " + strings.Repeat("*", n-4) + " *\n"    // prima e ultima  riga del quadrato dentro + fuori

	if n == 5 {

		return riga + "\n" + riga2 + strings.Repeat(riga3, n-4) + riga2 + riga
	}

	riga4 := "* *" + strings.Repeat(" ", n-6) + "* *\n" //seconda e penultima riga del quadrato dentro + fuori

	return riga + "\n" + riga2 + riga3 + strings.Repeat(riga4, n-6) + riga3 + riga2 + riga

}

func main() {

	if len(os.Args) == 1 {

		fmt.Println("too few arguments")
		return
	}

	grandezza, err := strconv.Atoi(os.Args[1])

	if err != nil {

		fmt.Println("required integer value")
		return
	}

	fmt.Println(NestedSquares(grandezza))

}

/*

Scrivere un programma 'squares.go' dotato di:

- una funzione
	func NestedSquares(n int) string
che costruisce e restituisce una stringa corrispondente ad una figura
formata da un quadrato di '*' di dimensione n (con n > 4), che racchiude
un quadrato di '*' di dimensione 'n - 4', separato da quello esterno da una
"cornice" di spazi di spessore unitario (si vedano gli esempi).
La stringa restituita NON deve terminare con '\n'.
Se n < 5, la funzione restituisce la stringa vuota

- una funzione main() che legge da linea di comando un intero
e produce su standard output la figura di quadrati della misura fornita.
Se non ci sono dati sulla linea di comando, il programma stampa il messaggio
"too few arguments" e termina.
Se il valore passato da linea di comando non è un intero, il programma stampa il messaggio
"required integer value" e termina.


Esempi
------
per n = 5
*****
*   *
* * *
*   *
*****

per n = 6
******
*    *
* ** *
* ** *
*    *
******

per n = 7
*******
*     *
* *** *
* * * *
* *** *
*     *
*******

per n = 8
********
*      *
* **** *
* *  * *
* *  * *
* **** *
*      *
********

per n = 9
*********
*       *
* ***** *
* *   * *
* *   * *
* *   * *
* ***** *
*       *
*********
etc..
*/
