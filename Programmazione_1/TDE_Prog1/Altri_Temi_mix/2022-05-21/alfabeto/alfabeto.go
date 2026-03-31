package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {

	dizionario := make(map[string]string)

	var stringa string

	lunghezza := len(os.Args[1])

	for _, i := range os.Args[1] {

		if dizionario[string(i)] == "" {

			if unicode.IsUpper(i) {

				fmt.Print(string(i+32), "? ")
			} else {
				fmt.Print(string(i), "? ")

			}

			fmt.Scan(&stringa)

			dizionario[string(i)] = stringa

			if unicode.IsUpper(i) {

				dizionario[string(i+32)] = stringa
			} else if unicode.IsUpper(i) == false {

				dizionario[string(i-32)] = stringa
			}

		}

	}

	for _, i := range os.Args[1] {

		fmt.Print(string(dizionario[string(i)]))

		if lunghezza > 1 {

			fmt.Print(" - ")
		}

		if lunghezza == 1 {

			fmt.Println()
		}

		lunghezza--

	}

}

/*

Alfabeto
--------

Scrivere un programma (il cui file si deve chiamare 'alfabeto.go') che fa la compitazione, cioè converte le parole in sequenze di stringhe corrispondenti ai nomi delle lettere di cui sono composte (a -> "a", b -> "bi", c -> "ci", d -> "di", e -> "e", f -> "effe", ..., z -> "zeta")

Il programma legge una parola (una stringa composta di sole lettere) da linea di comando, e per ogni lettera incontrata per la prima volta chiede all'utente di immettere il nome corrispondente (alimentando un "dizionario").
Alla fine stampa la sequenza delle stringhe che compongono la compitazione della parola passata per argomento, separate da " - ".

L'unico controllo richiesto sull'input è che ci sia almeno un argomento (parola).
Si assuma che le stringhe passate come argomenti contangano solo lettere e non altri caratteri.
Gli eventuali argomenti supplementari devono essere ignorati.
Il programma non distingue tra lettere maiuscole e minuscole.


Ad esempio, per la parola "Zappatrice", il programma stamperà:
zeta - a - pi - pi - a - ti - erre - i - ci - e

Esempio di esecuzione
---------------------

$ go run alfabeto.go sasso
s? esse
a? a
o? o
esse - a - esse - esse - o


*/
