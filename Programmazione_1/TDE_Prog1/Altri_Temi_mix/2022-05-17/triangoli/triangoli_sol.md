package main

import (
	"errors"
	"fmt"
	"math"
)

type Punto struct {
	x, y float64
}

type Triangolo struct {
	P1, P2, P3 Punto
}

func lunghezzeLati(A, B, C Punto) [3]float64 {

	var array [3]float64

	distanzaAB := math.Sqrt(((B.x - A.x) * (B.x - A.x)) + ((B.y - A.y) * (B.y - A.y)))
	distanzaBC := math.Sqrt(((C.x - B.x) * (C.x - B.x)) + ((C.y - B.y) * (C.y - B.y)))
	distanzaCA := math.Sqrt(((A.x - C.x) * (A.x - C.x)) + ((A.y - C.y) * (A.y - C.y)))

	array[0], array[1], array[2] = distanzaAB, distanzaBC, distanzaCA
	return array
}

func newTriangolo(A, B, C Punto) (Triangolo, error) {

	array := lunghezzeLati(A, B, C)

	a, b, c := array[0], array[1], array[2]

	var err error

	sommaAB, sommaBC, sommaCA := a+b, b+c, c+a

	if a < sommaBC && b < sommaCA && c < sommaAB {

		return Triangolo{A, B, C}, err
	}

	return Triangolo{}, errors.New("")
}

func tipoTriangolo(triangolo Triangolo) int {

	array := lunghezzeLati(triangolo.P1, triangolo.P2, triangolo.P3)

	if math.Abs(array[0]-array[1]) < 1e-6 && math.Abs(array[1]-array[2]) < 1e-6 {

		return 3
	}

	if math.Abs(array[0]-array[1]) < 1e-6 || math.Abs(array[1]-array[2]) < 1e-6 {

		return 2
	}

	return 0
}

func main() {

	var A, B, C Punto

	fmt.Scan(&A.x, &A.y, &B.x, &B.y, &C.x, &C.y)

	fmt.Println(lunghezzeLati(A, B, C))

	triangolo, err := newTriangolo(A, B, C)

	if err != nil {

		fmt.Println("Non è un triangolo")
		return
	}

	fmt.Println("triangolo ", triangolo)
	fmt.Println("tipo ", tipoTriangolo(triangolo))

}

/**

Esercizio "triangoli"
=====================

Scrivere un programma Go (il file deve chiamarsi 'triangoli.go') dotato di:

- una struttura Punto(x,y), dove x e y sono float64

- una struttura Triangolo(P1,P2,P3), dove P1, P2 e P3 sono di tipo Punto

- una funzione

	 lunghezzeLati(A, B, C Punto) [3]float64

  che restituisce una array di tre elementi corrispondenti, nell'ordine, alle distanze
 A-B, B-C e C-A

- una funzione

	newTriangolo(A, B, C Punto) (Triangolo, err)

  che restituisce il triangolo A, B, C se i punti A, B, C determinano
  effettivamente un triangolo (*), altrimenti restituisce un errore

- una funzione

	tipoTriangolo(triangolo Triangolo) int

  che passato un triangolo come parametro, restituisce il numero di
  lati di uguale (**) lunghezza, cioè:

	- 0 se il triangolo è scaleno

	- 2 se il triangolo è isoscele

	- 3 se il triangolo è equilatero

- una funzione

	main()

  che legge le coordinate x,y  (float64) di tre punti A, B, C (un punto per riga) da
  standard input ed emette su standard output le distanze A-B, B-C e C-A.
  Inoltre, se i punti non formano un triangolo, emette il messaggio di
  errore "non è un triangolo", altrimenti emette il triangolo e il suo tipo.

  (*) Disuguaglianza triangolare: In un triangolo la misura di ciascun lato è sempre
  strettamente minore della somma delle misure degli altri due lati.

  (**) Avvertenza: per verificare se due valori double a,b sono "uguali" si suggerisce
  di utilizzare la seguente approssimazione:
  a e b sono uguali se e solo se |a - b| < precisione
  dove precisione sia fissata a 1e-6 ("10 alla -6")

Esempi
------

1) dato come input:

0 -2
-2 0
0 0

l'output è:
[2.8284271247461903 2 2]
triangolo {{0 2} {2 0} {0 0}}
tipo 2

2) dato come input:

1 1
2 2
3 3

l'output è:
[1.4142135623730951 1.4142135623730951 2.8284271247461903]
non è un triangolo

*/
