package main

import (
	"fmt"
	"io"
	"math"
)

type Punto struct {
	x, y float64
}

func Distanza(p1, p2 Punto) float64 {

	return math.Sqrt(((p2.y - p1.y) * (p2.y - p1.y)) + ((p2.x - p1.x) * (p2.x - p1.x)))
}

func percorso(tappe []Punto) float64 {

	var distanza float64

	for i := 0; i < len(tappe)-1; i++ {

		metri := Distanza(tappe[i], tappe[i+1])

		distanza += metri
	}

	distanza += Distanza(tappe[0], tappe[len(tappe)-1])

	return distanza

}

func main() {

	var punto Punto

	slice := []Punto{}

	for {

		_, err := fmt.Scan(&punto.x, &punto.y)

		if err == io.EOF {

			break
		}

		slice = append(slice, punto)
	}

	fmt.Println("tappe:", slice)

	fmt.Println(percorso(slice))

}

/*
Percorso
--------

Scrivere un programma percorso.go che ha:
- una struttura Punto con due campi float64 x, y che
rappresentano le coordinate del punto nel piano cartesiano.

- una funzione
	distanza(p1,p2 Punto) float64
che, dati due Punti p1 e p2, restituisce la distanza tra p1 e p2.

- una funzione
	percorso(tappe []Punto) float64
che, data una sequenza di tappe, rappresentata da una slice di Punti,
restituisce la distanza da percorrere, partendo dalla prima tappa,
per tornare a questa, passando per tutte le altre tappe in ordine.

- una funzione
	main()
che legge da standard input una sequenza (non vuota) di valori numerici (floating point)
x1, y1, x2, y2, ..., xn, yn (per terminare la sequenza, dare invio e CTRL-D)
e stampa la sequenza di punti letti e la distanza totale da percorrere per andare
dal primo al secondo punto, dal secondo al terzo, ... e dall'ultimo al primo.

Nota : si puo` assumere, senza fare controlli, che la sequenza in input contenga
almento i dati di un punto e che i dati siano nella forma attesa (float).


Esempi
------

Input: 0 3 0 0 4 0
Output:
tappe: [{0 3} {0 0} {4 0}]
percorso 12

Input: 1.5 3.2 5.1 7.6
Output:
tappe: [{1.5 3.2} {5.1 7.6}]
percorso 11.370136322841516
*/
