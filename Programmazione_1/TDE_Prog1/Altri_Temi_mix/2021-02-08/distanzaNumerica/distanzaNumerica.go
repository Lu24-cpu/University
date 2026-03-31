package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var somma, corrente, precedente int

	if len(os.Args) == 1 || len(os.Args) == 2 {

		fmt.Println("NON OK")

		return
	}

	numeri := os.Args[1:]

	corrente, _ = strconv.Atoi(numeri[0])

	somma += corrente

	for i := 1; i < len(numeri); i++ {

		precedente = corrente

		corrente, _ = strconv.Atoi(numeri[i])

		somma += corrente

		if math.Abs(float64(corrente)-float64(precedente)) != float64(len(os.Args)-1) {

			fmt.Println("NON OK")
			return
		}

	}

	fmt.Println(somma)

}

/*
Scrivere un programma 'distanzaNumerica.go' che legge da linea di comando
una serie di numeri interi e verifica che la distanza numerica (differenza in valore assoluto)
fra due elementi consecutivi della sequenza sia sempre corrispondente
al numero di argomenti passati.
Se lo è, stampa la somma dei numeri; altrimenti stampa NON OK.
Anche nel caso la sequenza contenga nessun o un solo elemento, il programma stampa NON OK.

Esempi
------

$ go run distanzaNumerica.go 1 4 7
12

$ go run distanzaNumerica.go 10 7 7 7 8
NON OK

$ go run distanzaNumerica.go 10 5 0 -5 0
10

*/
