// testo esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num []int
	var max, min, i, sum int
	var media float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if i == 0 {
			max, min = n, n
		}
		num = append(num, n)
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
		sum += n
		i++
	}

	media = float64(sum) / float64(len(num))

	fmt.Printf("Min: %d\nMax: %d\nMedia: %.11f\n", min, max, media)
}

/*

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
