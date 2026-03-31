// testo dell'esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() {
	fmt.Println("+ (Aggiungi frasi al dizionario)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")
}

func main() {
	var testo []string
	Menu()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		riga := scanner.Text()

		parole := strings.Split(riga, " ")
		switch parole[0] {
		case "+":
			fmt.Println("Alimento dizionario")
		case "?":
			fmt.Println("Consulto dizionario")
		case "m":
			fmt.Println("Lunghezza minima e massima")
		case "p":
			fmt.Println("Stampa ordinata del dizionario")
		default:
			fmt.Println("Opzione non valida")
		}
	}
}

/*
Lessico v1
----------

Scrivere un programma 'lessico1.go' che si comporta come segue.

1) Stampa il seguente menu di opzioni:
  - (legge e memorizza)
    ? (numeri di riga in cui è comparsa la parola data)
    m (stampa le lunghezze min e max)
    p (stampa le parole memorizzate)

2) Legge da standard input un testo di un numero arbitrario di righe e termina quando riceve un "end of file" (cioè EOF, da tastiera pressione di 'CTRL-d' dopo aver dato <invio>).

Quando la riga letta **inizia** con:
- "+" : il programma stampa "alimento il dizionario"
- "?" : il programma stampa "consulto il dizionario"
- "m" : il programma stampa "lunghezza min e max"
- "p" : il programma stampa "stampo il dizionario ordinato"
- per qualsiasi altro carattere il programma stampa "opzione non valida"

Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

+ la befana ha il fazzoletto e la gonna rattoppata
? ha
n
+ ma quest'anno poverina
+ la befana è raffreddata!
? ha il
m
p

il programma produce il seguente output:

+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)
alimento il dizionario
consulto il dizionario
opzione non valida
alimento il dizionario
alimento il dizionario
consulto il dizionario
lunghezza min e max
stampo il dizionario ordinato
*/
