package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Vocabolario map[string]int

func (voc Vocabolario) String() {

	slice := []string{}

	for i := range voc {

		slice = append(slice, i)
	}

	sort.Strings(slice)

	for _, i := range slice {

		fmt.Println(i, ":", voc[i])

	}
}

func statistiche(voc Vocabolario) (int, int, float64) {

	var max int

	var count, somma float64

	var min int = math.MaxInt

	for s := range voc {

		count++

		if voc[s] > max {

			max = voc[s]
		} else if voc[s] < min {

			min = voc[s]
		}

		somma += float64(voc[s])

	}

	return min, max, somma / count
}

func selezione(voc Vocabolario, char rune) []string {

	slice := []string{}

	for i := range voc {

		if rune(i[0]) == char {

			slice = append(slice, i)

		}

	}

	sort.Strings(slice)

	return slice
}

func main() {

	if len(os.Args) != 3 {

		fmt.Println("indicare nome file e iniziale")
		return
	}

	voc := make(Vocabolario)

	file, err := os.Open(os.Args[1])

	runa := []rune(os.Args[2])

	if err != nil {

		fmt.Println("WProblema nell'apertura del file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		riga := scanner.Text()

		riga1 := strings.TrimSpace(riga)

		paroleRiga := strings.Split(riga1, " ")

		for _, i := range paroleRiga {

			if i != "" {

				voc[i]++
			}

		}

	}

	voc.String()

	fmt.Println(statistiche(voc))

	fmt.Println(selezione(voc, runa[0]))

}

/*

Analisi di testo (NOTA: CASE SENSITIVE O NO?)
================

Scrivere un programma in Go (il file deve chiamarsi 'analisiTesto.go') dotato di

- un tipo Vocabolario  in cui salvare le stringhe (sequenze di caratteri separate da white spaces) di un testo e, per ciascuna, il numero di volte che essa appare nel testo

- un metodo String per Vocabolario che stampa il Vocabolario in ordine lessicografico, un lemma per riga seguito da spazio,':', spazio e dalla sua frequenza. Ad esempio:
Lorem : 3
ever : 1

- una funzione

		statistiche(voc Vocabolario) (int, int, float64)

  che restituisce la frequenza più bassa, la frequenza più alta e la media delle frequenze delle stringhe del Vocabolario

- una funzione

		selezione(voc Vocabolario, char rune) []string

	che restituisce la lista delle stringhe contenute nel Vocabolario che iniziano con il carattere (runa) char

- una funzione

		main()

  che legge da linea di comando il nome di un file e un carattere (runa);
  legge il testo contenuto nel file e crea un Vocabolario delle stringhe contenute nel testo, ciascuna con la frequenza con cui appare nel testo;
  stampa il Vocabolario creato, i risultati restituiti dalla funzione statistiche e la lista di stringhe restituita dalla funzione selezione (vedi esempio).
  Se non viene passato un argomento per il nome del file e/o un carattere, il programma deve stampare il messaggio "indicare nome file e iniziale".
  Se si verifica un errore nell'apertura del file, il programma deve stampare l'errore restituito per la mancata apertura del file.

Esempio
-------
(sul file uno.input messo a disposizione)

$ ./analisiTesto uno.input s
Bëatrice : 1
Quant’ : 1
al : 1
bene : 1
che : 2
ch’era : 1
color, : 1
convenia : 1
da : 1
dentro : 1
di : 1
dov’ : 1
entra’mi, : 1
esser : 1
in : 1
io : 1
lucente : 1
lume : 1
l’atto : 1
ma : 1
meglio, : 1
non : 2
parvente! : 1
per : 3
quel : 1
quella : 1
scorge : 1
si : 1
sol : 1
sporge. : 1
subitamente : 1
suo : 1
sé : 1
sì : 2
tempo : 1
È : 1
1 3 1.1388888888888888
[scorge si sol sporge. subitamente suo sé sì]
*/
