// testo esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
)

func DrawPoint(c byte, k int) (s0 string) {
	for i := 0; i < k-1; i++ {
		s0 += " "
	}
	s0 += string(c)
	return
}

func DrawSegment(c byte, k, l int) (s1 string) {
	for i := 0; i < k-1; i++ {
		s1 += " "
	}
	for i := 0; i < l; i++ {
		s1 += string(c)
	}
	return
}

func main() {
	var c byte
	var s string
	ks, ls, car := os.Args[1], os.Args[2], os.Args[3]

	ki, _ := strconv.Atoi(ks)
	li, _ := strconv.Atoi(ls)

	for _, carattere := range car {
		c = byte(carattere)
	}

	for i := 0; i < li; i++ {
		s = DrawSegment(c, i+len(s)-1, ki)
		fmt.Println(s)
		for j := 0; j < ki-1; j++ {
			fmt.Println(DrawPoint(c, len(s)))
		}
	}
}

/*
draw
------

Scrivere un programma 'draw.go' dotato di:

- una funzione
	DrawPoint(c byte, k int) string
che restituisce una stringa formata da k spazi bianchi
seguiti dal carattere c

- una funzione
	DrawSegment(c byte, k, l int) string
che restituisce una stringa formata da k spazi bianchi
seguiti da l caratteri c

- una funzione
	main()
che legge come parametri da linea di comando (in quest'ordine)
due numeri interi l e n e un carattere c (byte),
e, se l e n sono > 0, produce su standard output una scala di n gradini
di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
altrimenti non fa niente.

Si può assumere che il programma venga lanciato con tre parametri
dei tipi attesi (non occorre cioè fare controlli).


Esempi
------

$ go run draw.go 3 1 x
xxx
  x
  x



$ go run draw.go 3 2 +
+++
  +
  +
  +++
    +
    +

*/
