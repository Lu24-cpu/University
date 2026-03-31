// Testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var count, n int
	var err error
	var aspettativa bool

	if len(os.Args) < 2 {
		fmt.Println("Nessun valore in ingresso")
		return
	}

	for i, elemento := range os.Args[1:] {
		n, err = strconv.Atoi(elemento)

		if err != nil {
			fmt.Println("elemento in posizione", i+1, "non valido")
			break
		}

		if i == 0 && n%2 == 0 {
			aspettativa = true
			continue
		}
		if i == 0 && n%2 == 1 {
			aspettativa = false
			continue
		}
		if aspettativa && n%2 == 1 {
			aspettativa = !aspettativa
		} else if !aspettativa && n%2 == 0 {
			aspettativa = !aspettativa
		} else {
			count = i + 1
			break
		}
	}

	if count == 0 && err == nil {
		fmt.Println("sequenza valida")
	} else {
		fmt.Println("Elemeneto in posizione", count, "non valido")
	}
}

/*

Scrivere un programma 'pariDispari.go' che legge una sequenza di valori interi da linea di comando
e controlla che si alternino valori pari e valori dispari.
In questo caso il programma stampa il messaggio "sequenza valida", altrimenti "elemento in posizione i non valido",
dove i è la posizione del primo  elemento (da sinistra) che non rispetta la regola
di alternanza o che non è un valore numerico.

In caso di mancanza di valori, il programma deve stampare "nessun valore in ingresso".

La sequenza può iniziare sia con un valore pari sia con uno dispari.

Si ricorda che lo zero è un numero pari.

Esempi
------

Argomenti:
	3 8 1 12
Output:
	sequenza valida

Argomenti:
	4
Output:
	sequenza valida

Argomenti:
	1 2 3 5
Output:
	elemento in posizione 4 non valido

Argomenti:
	1 2 3eqeqw 5
Output:
	elemento in posizione 3 non valido

Argomenti:

Output:
	nessun valore in ingresso

*/
