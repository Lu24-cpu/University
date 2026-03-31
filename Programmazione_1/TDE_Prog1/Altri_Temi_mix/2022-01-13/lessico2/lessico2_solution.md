package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func PrintMenu() {

	fmt.Println("+ (legge e memorizza)")
	fmt.Println("? (numeri di riga in cui è comparsa la parola data)")
	fmt.Println("m (stampa le lunghezze min e max)")
	fmt.Println("p (stampa le parole memorizzate)")
}

func PrintDict(m map[string][]int) {

	slice := []string{}

	for v := range m {

		slice = append(slice, v)

	}

	sort.Strings(slice)

	for _, v := range slice {

		fmt.Printf("%s : %d\n", v, m[v])

	}
}

func AddToDict(line string, n int, dict map[string][]int) {

	riga := strings.Split(line, " ")

	for i := 1; i < len(riga); i++ {

		dict[riga[i]] = append(dict[riga[i]], n)
	}

}

func Stats(dict map[string][]int) (int, int) {

	var max int

	min := math.MaxInt

	for chiave := range dict {

		if len(chiave) > max {

			max = len(chiave)

		}

		if len(chiave) < min {

			min = len(chiave)

		}

	}
	return min, max
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanLines)

	var c int

	var riga string

	slice := []string{}

	dict := make(map[string][]int)

	for scanner.Scan() {

		riga = scanner.Text()

		slice = append(slice, riga)

	}
	PrintMenu()

	for _, v := range slice {

		switch v[0] {

		case '+':
			fmt.Println("alimento il dizionario")
			c++
			AddToDict(v[1:], c, dict)

		case '?':

			fmt.Println("consulto il dizionario")
			fmt.Println("parola:", v[2:])
			fmt.Println("righe", dict[v[2:]])

		case 'm':
			fmt.Println("lunghezza min e max")

			fmt.Println(Stats(dict))

		case 'p':
			fmt.Println("stampo il dizionario ordinato")

			(PrintDict(dict))

		default:
			fmt.Println("opzione non valida")

		}
	}

}

/*
Lessico v2
----------

Copiare il programma 'lessico1.go' in un file 'lessico2.go'

Il programma 'lessico2.go' alimenta un "dizionario" con le parole del testo e le righe (i numeri di riga) in cui ciascuna parola compare; inoltre permette di consultare e stampare il "dizionario".

In particolare, si richiede di dotare il programma 'lessico2.go' delle seguenti funzioni, che andranno usate nel main:

- func PrintMenu()
che stampa il menu delle opzioni, cioè il seguente menu:
+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)

- func PrintDict(m map[string][]int)
che, data una mappa m, stampa in ordine lessicografico, una per riga, le chiavi della mappa, ciascuna seguita, sulla stessa riga, da uno spazio, un ':', uno spazio e dal suo valore (vedi Esempio sotto)

- AddToDict(line string, n int, dict map[string][]int)
(vedi sotto, opzione '+') che riceve la riga letta che inizia con '+' (line) insieme alla sua posizione nel testo (numero di riga n) e gestisce l'aggiornamento del "dizionario".
Si può assumere, senza fare controlli, che la riga abbia il formato atteso, cioè: '+' seguito da uno spazio e poi da un testo

- Stats(dict map[string][]int) (int, int)
che restiluisce la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti in dict

Aggiungere inoltre al programma 'lessico2.go', per ogni opzione del menu, dopo le istruzioni di stampa della versione 1, le funzionalità specificate qui sotto.

Quando la riga letta inizia con:
- "+" (alimenta dizionario): il programma usa il rimanente della riga e memorizza in un "dizionario" le parole (stringhe separate da white space) che la costituiscono e il numero di riga. La prima riga letta è la numero 1;
- "?" (consulta dizionario): il programma usa il rimanente della riga come stringa per la consultazione del dizionario e stampa i numeri di riga associati a tale stringa;
- "m" (determina min e max): il programma stampa la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti nel dizionario
- "p" (print): il programma stampa le parole presenti nel  "dizionario", una per riga e in ordine lessicografico, ciascuna seguita dalla lista dei numeri di riga in cui la parola è comparsa (per il formato, vedi Esempio sotto).


Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

+ aa bbbb ccc
+ aa
q
p
m
? aa
? aa bbbb

il programma produce il seguente output:

+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)
alimento il dizionario
alimento il dizionario
opzione non valida
stampo il dizionario ordinato
aa : [1 2]
bbbb : [1]
ccc : [1]
lunghezza min e max
2 4
consulto il dizionario
parola: aa
righe [1 2]
consulto il dizionario
parola: aa bbbb
righe []


*/
