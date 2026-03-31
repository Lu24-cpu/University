package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("+ (Legge una riga e memorizza le parole)\n")
	fmt.Printf("? (Legge una parola e indica le righe che la contengono)\n")
	fmt.Printf("p (Stampa le parole)\n")

	dizionario := make(map[string][]int)

	var line int = 1

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		riga := strings.Split(scanner.Text(), " ")

		switch riga[0] {

		case "+":
			for _, i := range riga[1:] {

				dizionario[i] = append(dizionario[i], line)
			}
			line++
		case "?":

			fmt.Println("parola:", riga[1:])

			fmt.Print("righe:")
			for _, i := range riga[1:] {

				fmt.Println(dizionario[i])
			}

		case "p":
			fmt.Println(dizionario)
		}

	}

}

/*
Lessico
-------

Scrivere un programma lessico.go che
- stampa il seguente menu di opzioni
	+ (Legge una riga e memorizza le parole)
	? (Legge una parola e indica le righe che la contengono)
	p (Stampa le parole)

- legge stringhe da standard input
- il programma termina quando riceve un "end of file" (cioè EOF, pressione di 'CTRL-d')

Se la stringa inizia con:
- "+" (alimenta dizionario): il programma usa il rimanente della riga e memorizza in un "dizionario" le parole che la costituiscono;
- "?" (consulta dizionario): il programma usa il rimanente della riga e stampa i numeri di riga del dizionario 	in cui è comparsa la stringa;
- "p" (print): il programma stampa le parole presenti nel  "dizionario", con l’elenco dei numeri di riga in cui compaiono;


Esempio
-------

+ (Legge una riga e memorizza le parole)
? (Legge una parola e indica le riga che la contengono)
p (Stampa le parole)
+ la befana ha il fazzoletto e la gonna rattoppata
p
map[rattoppata:[1] la:[1 1] befana:[1] ha:[1] il:[1] fazzoletto:[1] e:[1] gonna:[1]]
+ ma quest'anno poverina la befana è raffreddata
? la
parola: la
righe [1 1 2]
? befana
parola: befana
righe [1 2]
? il
parola: il
righe [1]
p
map[raffreddata:[2] la:[1 1 2] il:[1] fazzoletto:[1] ma:[2] quest'anno:[2] poverina:[2] è:[2] befana:[1 2] ha:[1] e:[1] gonna:[1] rattoppata:[1]]

*/
