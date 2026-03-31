package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdout)

	var count float64

	var max, somma float64

	var min float64 = math.MaxFloat64

	for scanner.Scan() {

		count++

		numero, _ := (strconv.ParseFloat(scanner.Text(), 64))

		somma += numero

		if numero > max {

			max = numero
		} else if numero < min {

			min = numero
		}

	}

	fmt.Println("min: ", int(min))
	fmt.Println("max: ", int(max))
	fmt.Println("miedia ", somma/count)

}

/**

Esercizio filtro
================

Scrivere un programma in Go (il file deve chiamarsi 'filtro.go') che,
data una serie (di lunghezza arbitraria) di numeri interi da standard input
(la serie contiene almeno un elemento, e termina con CTRL-D su una riga vuota),
calcola ed emette su standard output:

- il valore minimo (int)

- il valore massimo (int)

- la media (in tipo float64, quindi non troncata)

dei valori letti. Per il formato si vedano gli esempi qui sotto.

Si fornisce un file "esempioLungo.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -3556
max: 9292
media: 2871.39393939394

E un file "esempioBreve.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -26
max: 77
media: 27.75

Nota bene: la verifica della correttezza del filtro avviene mediante controllo automatico
(uno script) ergo attenersi strettamente all'output specificato sopra,
pena una possibile esclusione dalla prova d'esame.

*/
