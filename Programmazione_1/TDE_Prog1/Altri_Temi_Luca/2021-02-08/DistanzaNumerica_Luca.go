// testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
)

func Modulo(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var now, prec, moduleprec, modulenow, sum int
	var flag bool = true

	nums := os.Args[1:]

	if len(nums) == 0 || len(nums) == 1 {
		fmt.Println("NON OK")
	}
	for _, value := range nums {
		if value == nums[0] {
			now, _ = strconv.Atoi(value)
			sum += now
			continue
		}
		if value == nums[1] {
			modulenow = Modulo(prec - now)
		}
		prec = now
		now, _ = strconv.Atoi(value)
		sum += now
		modulenow = Modulo(prec - now)

		if moduleprec == modulenow {
			flag = true
		} else if value == nums[1] || value == nums[0] {
			flag = false
		}
		moduleprec = modulenow
	}

	if !flag {
		fmt.Println("NON OK")
	} else {
		fmt.Println(sum)
	}

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
