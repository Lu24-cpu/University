// testo esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckInput(s string) int {
	n, err := strconv.Atoi(s)

	if n < -1 || err != nil {
		return -2
	} else {
		return n
	}
}

func leastDigit(n int) int {
	s := strconv.Itoa(n)

	n1, _ := strconv.Atoi(string(s[len(s)-1]))
	return n1
}

func main() {
	var num []int
	var flag bool
	lastD := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserire solo numeri positivi")
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		for _, el := range parts {
			n := CheckInput(el)

			if n > -1 {
				num = append(num, n)
				lastD[leastDigit(n)]++
			} else if n == -1 {
				num = append(num, n)
				flag = true
			}
		}
		if flag {
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, "-", lastD[i])
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
