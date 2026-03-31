// Testo dell'esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var righe []string
	nrighe := os.Args[1]
	nomefile := os.Args[2]

	NRighe, _ := strconv.Atoi(nrighe)
	file, _ := os.Open(nomefile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		righe = append(righe, scanner.Text())
	}

	if len(righe) < NRighe {
		for _, riga := range righe {
			fmt.Println(riga)
		}
	} else {
		for i := len(righe) - NRighe; i < len(righe); i++ {
			fmt.Println(righe[i])
		}
	}

}

/*
# TESTO

Realizzare un programma Go (nome file 'tail.go') che implementi un semplice 'tail', comando Unix che estrae le ultime N righe di un file di testo.

Il programma deve prendere due parametri da linea di comando (in quest'ordine):
- numero N di linee da estrarre
- nome del file da elaborare

e stampare su standard output le ultime N righe del file. Se il file ha meno di N righe, il programma stamperà tutte quelle che ci sono.

Nota bene: NON va implementato invocando da Go il comando 'tail' di sistema.

## Esempio esecuzione

(presuppone il 'tail.go' già compilato, 'uno.input' è presente in questa directory)

$ ./tail 3 uno.input
remaining essentially unchanged. It was popularised in the 1960s with the release
of Letraset sheets containing Lorem Ipsum passages, and more
recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

*/
