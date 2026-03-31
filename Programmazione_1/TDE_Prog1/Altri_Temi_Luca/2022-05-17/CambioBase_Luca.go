// testo esercizio a fine codice

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func TenToB(n, b int) (s string, err error) {
	if b < 2 || b > 16 {
		err = errors.New("Base non valida")
		return
	}
	values := "0123456789ABCDE"[:b]
	var resto int

	err = nil
	for n != 0 {
		resto = n % b
		s += string(values[resto])
		if resto == 0 {
			n /= b
		} else {
			n = (n - 1) / b
		}
	}
	return
}

func main() {
	n, b := os.Args[1], os.Args[2]

	if len(os.Args) > 3 {
		fmt.Println("Numero parametri non valido")
	}
	nint, err1 := strconv.Atoi(n)
	bint, err2 := strconv.Atoi(b)

	if err1 != nil {
		fmt.Println("parametro non valido")
	}
	if err2 != nil {
		fmt.Println("base non valida")
	}

	nbaseb, errn := TenToB(nint, bint)

	if errn != nil {
		fmt.Println(errn)
	}

	for i := len(nbaseb) - 1; i >= 0; i-- {
		fmt.Printf("%c", nbaseb[i])
	}
	fmt.Println()
}

/*

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
