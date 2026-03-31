package main

import (
	"fmt"
	"strconv"
)

func checkInput(s string) int {

	numero, err := strconv.Atoi(s)

	if numero >= -1 && err == nil {

		return numero

	}

	return -2
}

func leastDigit(n int) int {

	return n % 10
}

func main() {

	var numero string

	mappa := make(map[int]int)

	for {

		fmt.Scan(&numero)

		n, _ := strconv.Atoi(numero)

		if numero == "-1" {

			break
		}

		if checkInput(numero) != -2 {

			mappa[leastDigit((n))]++

		}
	}

	for i := 0; i < 10; i++ {

		fmt.Println(i, "-", mappa[i])
	}

}

/*
Scrivere un programma in Go (il file deve chiamarsi 'leastDigit.go') che, data una serie di numeri interi non negativi da standard input, terminata da -1, conta quanti di questi numeri (escluso il -1) hanno come cifra meno significativa (cifra delle unità) la cifra 0, quanti la cifra 1, quanti la cifra 2, ...., quanti la cifra 9.
Il programma ignora, senza terminare, eventuali dati (tranne -1 che è il terminatore dell'input) che non siano interi non negativi.
Il programma stampa infine quanto calcolato (vedi esempio sotto).

Il programma deve essere dotato di:

- una funzione
	checkInput(s string) int
che restituisce -2 se la stringa s non rappresenta un intero >= -1 (se quindi s rappresenta un int < -1 o un non int), e altrimenti restituisce l'intero corrispondente a s.

- una funzione
	leastDigit(n int) int
che restituisce la cifra meno significativa (quella più a destra, la cifra delle unità) del numero n.


Esempio
-------
dato il seguente input

123
456 789 ciao uno due
-35
66 88 99
-2
-1

il programma stampa
0 - 0
1 - 0
2 - 0
3 - 1
4 - 0
5 - 0
6 - 2
7 - 0
8 - 1
9 - 2

*/
