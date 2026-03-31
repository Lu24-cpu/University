// Testo a fine esercizio
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cancella(lista []string) []string {
	for i, riga := range lista {
		parts := strings.Split(riga, " ")
		for j, elemento := range parts {
			n, err := strconv.Atoi(elemento)
			if err == nil {
				if len(parts)-j > n {
					parts = append(parts[:j], parts[j+n+1:]...)
				} else {
					parts = parts[:j]
				}
			}
		}
		lista[i] = strings.Join(parts, " ")
	}
	return lista
}

func main() {
	var lista []string

	file, err := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err == nil {
		for scanner.Scan() {
			lista = append(lista, scanner.Text())
		}
	} else {
		fmt.Println("errore: ", err)
	}

	fmt.Println("Lista originale:")
	for _, riga := range lista {
		fmt.Println(riga)
	}

	fmt.Println("Lista modificata:")
	for _, riga := range Cancella(lista) {
		fmt.Println(riga)
	}
}

/*
Cancellazioni
=============

Scrivere un programma 'cancellazioni.go' dotato di:

- una funzione

 	func cancella(lista []string) []string

  che, per ogni numero n >0 (intero) presente in lista,
  cancella il numero stesso e gli n elementi successivi della lista
  (o quelli che ci sono per arrivare alla fine della lista)
  e restituisce la nuova lista così prodotta;

- una funzione main() che legge da file (il cui nome viene passato
  come parametro su linea di comando) un testo di parole (sequenze di caratteri separate da whitespace),
  tra cui anche numeri interi non negativi, stampa la lista delle parole lette e poi
  la nuova lista ottenuta cancellando, per ogni numero n presente nella lista originale,
  il numero stesso e gli n elementi successivi (vedi sopra).

Se il programma viene lanciato con un numero di argomenti diverso da 1,
deve terminare stampando "Fornire 1 nome di file".
Se invece il file non esiste, il programma deve terminare stampando "File non accessibile".

Il file può contenere un numero arbitrario di parole e numeri, disposti su un numero arbitrario di righe di testo, senza vincoli sul numero e tipo di caratteri whitespace che separano parole e numeri e sul numero di cifre di cui sono composti i numeri.

Si può assumere che il file non contenga numeri negativi, non occorre fare questo controllo.

Esempi di esecuzione
--------------------

$ ./cancellazioni uno.input
[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci]
[uno due cinque sette]

$ ./cancellazioni due.input
[uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]
[uno due cinque sette nove dieci]

$ ./cancellazioni quattro.input
[uno 10 due tre quattro cinque sei sette otto nove dieci undici dodici]
[uno dodici]

$ ./cancellazioni
Fornire 1 nome di file

$ ./cancellazioni  FileNonEsistente.txt
File non accessibile

*/
