// Testo Esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type PUNTO struct {
	x float64
	y float64
}

func Distanta(p1, p2 PUNTO) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func Percorso(tappe []PUNTO) (p2 float64) {
	for i := 1; i <= len(tappe)-1; i++ {
		p2 += Distanta(tappe[i], tappe[i-1])
	}
	p2 += Distanta(tappe[len(tappe)-1], tappe[0])
	return
}

func main() {
	var tappa PUNTO

	fmt.Println("Inserire più coppie di punti sulla stessa riga")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var tappe []PUNTO
		parts := strings.Split(scanner.Text(), " ")

		for i := 0; i < len(parts); i++ {
			tappa.x, _ = strconv.ParseFloat(parts[i], 64)
			tappa.y, _ = strconv.ParseFloat(parts[i+1], 64)
			tappe = append(tappe, tappa)
			i++
		}
		fmt.Println("tappe: ", tappe)
		fmt.Println("Percorso", Percorso(tappe))
	}

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
