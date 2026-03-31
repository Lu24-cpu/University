// Testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"unicode"
)

type MAPPA map[rune]string

func main() {
	var pronuncia string
	dizionario := make(MAPPA)
	parole := os.Args[1]

	parola := []rune(parole)
	for _, lettera := range parola {
		lettera = unicode.ToLower(lettera)
		if dizionario[lettera] == "" {
			fmt.Printf("%c? ", lettera)
			fmt.Scan(&pronuncia)
			dizionario[lettera] = pronuncia
		}
	}

	for i, lettera := range parola {
		lettera = unicode.ToLower(lettera)
		fmt.Print(dizionario[lettera])
		if i < len(parola)-1 {
			fmt.Print(" - ")
		}
	}
	fmt.Println()
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
