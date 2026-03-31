// testo esercizio a fine codice

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type PUNTO struct {
	x, y float64
}

type TRIANGOLO struct {
	p1, p2, p3 PUNTO
}

const PRECISIONE = 1e-6

func LunghezzeLati(A, B, C PUNTO) (l [3]float64) {
	l[0] = math.Sqrt((A.x-B.x)*(A.x-B.x) + (A.y-B.y)*(A.y-B.y))
	l[1] = math.Sqrt((C.x-B.x)*(C.x-B.x) + (C.y-B.y)*(C.y-B.y))
	l[2] = math.Sqrt((A.x-C.x)*(A.x-C.x) + (A.y-C.y)*(A.y-C.y))

	return
}

func NewTirangolo(A, B, C PUNTO) (t TRIANGOLO, err error) {
	l := LunghezzeLati(A, B, C)

	if l[0]+l[1] < l[2] || l[1]+l[2] < l[0] || l[0]+l[2] < l[1] {
		err = errors.New("non è un triangolo")
		return
	}

	t.p1, t.p2, t.p3 = A, B, C
	return
}

func TipoTriangolo(t TRIANGOLO) int {
	l := LunghezzeLati(t.p1, t.p2, t.p3)

	if math.Abs(l[0]-l[1]) > PRECISIONE && math.Abs(l[0]-l[2]) > PRECISIONE && math.Abs(l[2]-l[1]) > PRECISIONE {
		return 0
	} else if (math.Abs(l[0]-l[1]) < PRECISIONE && math.Abs(l[2]-l[1]) > PRECISIONE && math.Abs(l[0]-l[2]) > PRECISIONE) || (math.Abs(l[2]-l[1]) < PRECISIONE && math.Abs(l[2]-l[0]) > PRECISIONE && math.Abs(l[0]-l[1]) > PRECISIONE) || (math.Abs(l[0]-l[2]) < PRECISIONE && math.Abs(l[2]-l[1]) > PRECISIONE && math.Abs(l[0]-l[1]) > PRECISIONE) {
		return 2
	} else {
		return 3
	}
}

func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

func main() {
	var punto [3]PUNTO
	var i int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		punti := scanner.Text()
		parts := strings.Split(punti, " ")

		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		punto[i].x = abs(x)
		punto[i].y = abs(y)
		i++
	}

	fmt.Println(LunghezzeLati(punto[0], punto[1], punto[2]))
	t, err := NewTirangolo(punto[0], punto[1], punto[2])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("triangolo", t)
		fmt.Println("tipo", TipoTriangolo(t))
	}

}

/*

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
