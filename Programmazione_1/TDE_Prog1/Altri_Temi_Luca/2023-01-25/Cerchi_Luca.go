// Testo dell'esercizio in fondo all'esercizio

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var MIND = 1e-6

type CERCHIO struct {
	nome         string
	x, y, raggio float64
}

func newCircle(s string) (cerchio CERCHIO) {
	parts := strings.Split(s, " ")
	cerchio.nome = parts[0]
	cerchio.x, _ = strconv.ParseFloat(parts[1], 64)
	cerchio.y, _ = strconv.ParseFloat(parts[2], 64)
	cerchio.raggio, _ = strconv.ParseFloat(parts[3], 64)

	return
}

func distanzaPunti(x1, y1, x2, y2 float64) (dis float64) {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func abs(n float64) float64 {
	if n > 0 {
		return n
	}
	return -n
}

func Tocca(cerchio1, cerchio2 CERCHIO) (tan bool) {
	dis := distanzaPunti(cerchio1.x, cerchio1.y, cerchio2.x, cerchio2.y)

	sumragg := cerchio1.raggio + cerchio2.raggio
	if dis > sumragg-MIND && dis < sumragg+MIND {
		return true
	}

	distanza := abs(cerchio2.raggio - cerchio1.raggio)
	if dis > distanza-MIND && dis < distanza+MIND {
		return true
	}

	return false
}

func Maggiore(nuovo, vecchio CERCHIO) bool {
	return nuovo.raggio > vecchio.raggio
}

func main() {
	var nuovo, vecchio CERCHIO
	scanner := bufio.NewScanner(os.Stdin)

	nuovo = CERCHIO{"", 0.0, 0.0, 0.0}
	for scanner.Scan() {
		fmt.Print("> ")
		riga := scanner.Text()
		nuovo, vecchio = newCircle(riga), nuovo
		tangente := Tocca(vecchio, nuovo)

		if tangente {
			if Maggiore(nuovo, vecchio) {
				fmt.Println(nuovo, "tangente, maggiore")
			} else {
				fmt.Println(nuovo, "tangente, minore o uguale")
			}
		} else {
			if Maggiore(nuovo, vecchio) {
				fmt.Println(nuovo, "non tangente, maggiore")
			} else {
				fmt.Println(nuovo, "non tangente, minore o uguale")
			}
		}
	}
}

/**
N.B. leggere il file README.txt per istruzioni di compilazione, test e consegna

=== Cerchi ===

Scrivere un programma 'cerchi.go' dotato di:

- una struttura 'Cerchio' con campi:
	nome string
	x,y,raggio float64
		(dove x e y sono le coordinate del centro)
  dichiarati in quest'ordine

- una funzione 'newCircle(descr string) Cerchio'
	che data una stringa che descrive un cerchio
	(nome, coordinate x e y del centro, raggio , in quest'ordine e separati da spazi)
	restituisce il cerchio corrispondente

- una funzione 'distanzaPunti(x1,y1,x2,y2 float64) float64
	che restituisce la distanza tra i punti di coordinate (x1,y1) e (x2,y2)

- una funzione 'tocca(cerc1, cerc2 Cerchio) bool'
	che restituisce true se i due cerchi sono tangenti
	(internamente o esternamente); false altrimenti.
	Trattandosi di valori float, consideriamo una distanza
	trascurabile se è inferiore a 10^-6 (1e-6)

- una funzione maggiore(cerc1, cerc2 Cerchio) bool
	che restituisce true se cerc1 è più grande di cerc2;
	false altrimenti.

- una funzione main()
	che legge da standard input una sequenza (terminata da ctrl D)
	di righe che descrivono cerchi, del tipo:

uno 10.3 12.7 45.8
due 1.3 2.9 5.8
pippo 7.3 22.5 6.6

	ciascuna contenente nome, coordinate del centro e raggio di un
	cerchio, in quest'ordine.

Il programma crea per ciascuna riga il cerchio corrispondente, lo stampa.
Inoltre stampa se il cerchio, rispetto a quello precedente, è tangente o no e maggiore o no.
Il primo confronto è fatto rispetto al cerchio "zero" ("", 0, 0, 0).
Non sono richiesti controlli sulla correttezza dei dati in input, potete assumere che siano sempre nell'ordine e del tipo previsto.




Esempio
-------
(vengono marcate con '>' le righe digitate da tastiera,
le altre sono l'output del programma)

>uno 0.5 0 2.5
{uno 0.5 0 2.5} non tangente, maggiore
>due 0 0 3
{due 0 0 3} tangente, maggiore
>tre 10.2 -8.4 1.5
{tre 10.2 -8.4 1.5} non tangente, minore o uguale

*/
