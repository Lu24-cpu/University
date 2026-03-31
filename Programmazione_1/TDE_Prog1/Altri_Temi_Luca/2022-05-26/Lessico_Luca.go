// Testo a fine esercizio

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MAPPA map[string][]int

func main() {
	var righe []string
	var i int
	mappa := make(MAPPA)
	fmt.Println("+ (Legge una riga e memorizza le parole)\n? (Legge una parola e indica le righe che la contengono)\np (Stampa le righe)")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		parts := strings.Split(riga, " ")

		if string(riga[0]) == "+" {
			i++
			righe = append(righe, scanner.Text()[2:])
			for _, parola := range parts {
				if parola == parts[0] {
					continue
				}
				mappa[parola] = append(mappa[parola], i)
			}
		}
		if riga[0] == '?' {
			fmt.Println("parola:", parts[1])
			fmt.Println("Righe:", mappa[parts[1]])

		}
		if riga[0] == 'p' {
			fmt.Println(mappa)
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
