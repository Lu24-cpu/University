// testo esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type VOCABOLARIO map[string]int

func (v VOCABOLARIO) String() {
	var words []string
	for key := range v {
		words = append(words, key)
	}
	sort.Strings(words)

	for _, parola := range words {
		fmt.Println(parola, ":", v[parola])
	}
}

func Statistiche(voc VOCABOLARIO) (min, max int, media float64) {
	var i, sum int
	for _, value := range voc {
		if i == 0 {
			min, max = value, value
		}
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
		sum += value
		i++
	}
	media = float64(sum) / float64(len(voc))
	return
}

func Selezione(voc VOCABOLARIO, char rune) (parole []string) {
	for key := range voc {
		parola := []rune(key)

		if len(parola) != 0 {
			if parola[0] == char {
				parole = append(parole, key)
			}
		}
	}
	return
}

func main() {
	var testo []string
	voc := make(VOCABOLARIO)

	if len(os.Args) < 3 {
		fmt.Println("indicare nome file e iniziale")
		return
	}

	filename, char := os.Args[1], os.Args[2]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		riga := scanner.Text()

		if riga == "" {
			continue
		}

		riga = strings.TrimSpace(riga)
		parole := strings.Split(riga, " ")

		testo = append(testo, parole...)
		for _, parola := range parole {
			if len(parola) != 0 {
				voc[parola]++
			}
		}
	}

	sort.Strings(testo)

	voc.String()
	fmt.Println(Statistiche(voc))
	charune := []rune(char)
	fmt.Println(Selezione(voc, charune[0]))

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
