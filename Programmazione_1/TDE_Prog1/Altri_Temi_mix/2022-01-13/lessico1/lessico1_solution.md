package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("+ (legge e memorizza)")
	fmt.Println("? (numeri di riga in cui è comparsa la parola data)")
	fmt.Println("m (stampa le lunghezze min e max)")
	fmt.Println("p (stampa le parole memorizzate)")

	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanLines)

	slice := []byte{}

	for scanner.Scan() {

		slice = append(slice, scanner.Text()[0])

	}

	for _, carattere := range slice {

		switch carattere {

		case '+':
			fmt.Println("alimento il dizionario")

		case '?':
			fmt.Println("consulto il dizionario")

		case 'm':
			fmt.Println("lunghezza min e max")

		case 'p':
			fmt.Println("stampo il dizionario ordinato")

		default:
			fmt.Println("opzione non valida")

		}
	}

}

/*
Lessico v1
----------

Scrivere un programma 'lessico1.go' che si comporta come segue.

1) Stampa il seguente menu di opzioni:
  + (legge e memorizza)
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
