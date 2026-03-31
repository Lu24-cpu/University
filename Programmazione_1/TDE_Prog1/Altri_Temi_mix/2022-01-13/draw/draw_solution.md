package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numero1, _ := strconv.Atoi(os.Args[1])
	numero2, _ := strconv.Atoi(os.Args[2])

	var spaziov int = numero1 - 1

	var spazio int

	carattere := []byte(os.Args[3])

	if numero1 > 0 && numero2 > 0 {

		for i := 0; i < numero2; i++ {

			fmt.Println(DrawSegment(byte(carattere[0]), spazio, numero1))

			fmt.Println(DrawPoint(byte(carattere[0]), spaziov))
			fmt.Println(DrawPoint(byte(carattere[0]), spaziov))

			spazio += numero1 - 1

			spaziov += numero1 - 1

		}
	}

}

func DrawPoint(c byte, k int) string {

	return strings.Repeat(" ", k) + string(c)
}

func DrawSegment(c byte, k, l int) string {

	return strings.Repeat(" ", k) + strings.Repeat(string(c), l)
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
