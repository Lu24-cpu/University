package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func tenToB(n, b int) (string, error) {

	slice := []string{}

	stringa := []rune("0123456789ABCDEF")

	var err error

	var numeroConvertito string

	var cifra int

	if b < 2 || b > 16 {

		err = errors.New("")

		return "", err
	}

	for {

		cifra = (n % b)

		if cifra == 0 && b != 2 {

			break
		}

		if b == 2 && n/b == 0 {

			slice = append(slice, string(stringa[cifra]))
			break
		}

		slice = append(slice, string(stringa[cifra]))

		n = n / b

	}

	for i := len(slice) - 1; i >= 0; i-- {

		numeroConvertito += slice[i]

	}

	return numeroConvertito, err

}

func main() {

	if len(os.Args) != 3 {

		fmt.Println("numero parametri non valido")
		return
	}

	n, err1 := strconv.Atoi(os.Args[1])
	b, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {

		fmt.Println("parametro non valido")
		return
	}

	if b < 2 || b > 16 {

		fmt.Println("base non valida")
		return
	}

	numeroCambiato, _ := tenToB(n, b)

	fmt.Println(numeroCambiato)

}

/**

Esercizio "cambio base"
=======================

Scrivere un programma Go (il file deve chiamarsi 'cambiobase.go') dotato di una funzione

- func tenToB(n, b int) (string, err)

che restituisce la rappresentazione in base b (b è il secondo parametro)
di un numero n (intero non negativo, in base 10) passato come primo parametro
e un valore di errore, diverso da nil se la base b non è compresa tra 2 e 16;

E di una funzione

- func main()

che legge due numeri interi, n e b, da linea di comando, ed emette
su standard output n espresso in base b.

Se il numero di parametri è diverso da due, se i parametri non sono
numeri interi, o se la base non è tra 2 e 16,
il programma deve stampare il messaggio corrispondente:
- "numero parametri non valido"
- "parametro non valido"
- "base non valida"

Suggerimento. Le cifre del risultato della conversione si ottengono prendendo
i resti della divisione di n per b, dividendo n per b fino a che il risultato
della divisione diventa 0.
I resti così ottenuti, in ordine inverso, formeranno il numero convertito.

Consideriamo ad esempio 6 in base 2

resto divisione 6 per 2 -> 0
quoziente divisione 6 per 2 -> 3
resto divisione 3 per 2 -> 1
quoziente divisione 3 per 2 -> 1
resto divisione 1 per 2 -> 1
quoziente divisione 1 per 2 -> 0 (stop)

Prendendo i resti in ordine inverso si ottiene il risultato: 110

Attenzione: per rappresentare numeri in base > 10 (e fino a base 16)
si devono usare anche le lettere A,B,C,D,E,F.
A tale scopo può essere conveniente usare la stringa "0123456789ABCDEF".

Esempi
------

$ go run cambiobase.go 13 3
111

$ go run cambiobase.go 4 20
base non valida

$ go run cambiobase.go 123 7
234

$ go run cambiobase.go 785 iuy
parametro non valido

$ go run cambiobase.go 140 13
AA

$ go run cambiobase.go 171 16
AB

$ go run cambiobase.go 123 2 4
numero parametri non valido

*/
