package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Appuntamento struct {
	giorno, orainizio, oraFine int
}

func NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool) {

	if gg >= 1 && gg <= 366 && h1 <= h2 && h1 >= 0 && h1 <= 24 && h2 >= 0 && h2 <= 24 {

		return Appuntamento{gg, h1, h2}, true

	}

	return Appuntamento{}, false
}

func CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool {

	if (app1.giorno != app2.giorno) || app1.giorno == app2.giorno && ((app1.oraFine <= app2.orainizio) || (app2.oraFine <= app2.orainizio)) {

		return false
	}

	return true
}

func AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool {

	for _, i := range *agenda {

		if CheckSovrapposizioneAppuntamenti(app, i) {

			return false
		}

	}

	*agenda = append(*agenda, app)

	return true
}

func AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento {

	var slice []Appuntamento

	for _, i := range agenda {

		if i.giorno == gg {

			slice = append(slice, i)
		}

	}

	return slice
}

func main() {

	agenda := []Appuntamento{}

	var pagenda *[]Appuntamento = &agenda

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		riga := scanner.Text()

		rigaS := strings.Split(riga, " ")

		gg, _ := strconv.Atoi(rigaS[0])
		h1, _ := strconv.Atoi(rigaS[1])
		h2, _ := strconv.Atoi(rigaS[2])

		AggiungiAppuntamento(Appuntamento{gg, h1, h2}, pagenda)
	}

	fmt.Println(agenda)

}

/*
Appuntamenti
------------

Scrivere un programma (il cui file deve chiamarsi 'appuntamenti.go') dotato di:

- un tipo Appuntamento con tre campi int (giorno, oraInizio, oraFine, dichiarati in quest'ordine) che rappresentano un giorno dell'anno, e l'ora di inizio e di fine  dell'appuntamento.
Si considerano per semplicità solo appuntamenti che iniziano e finiscono nella stessa giornata e a ore precise (intere).

- una funzione
	NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool)
che, se i parametri sono validi (giorno tra 1 e 366, ora inizio precedente o uguale a ora di fine, entrambe nell'intervallo [0,24])
crea un appuntamento corrispondente a questi dati e lo restituisce insieme al valore 'true',
altrimenti restituisce un Appuntamento "zero-value" e 'false'.

- una funzione
    CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool
che, dati due appuntamenti che si assume siano validi, restituisce 'true' se si sovrappongono (anche parzialmente), 'false' altrimenti.

- una funzione
	AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool
che, se l'appuntamento app non si sovrappone con NESSUNO degli altri già presenti in agenda, lo aggiunge alla stessa restituendo 'true', altrimenti non fa nulla e restituisce 'false'

- una funzione
    AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento
che, dati un giorno e una lista di appuntamenti, restituisce gli appuntamenti del giorno

- una funzione main() che crea un'agenda vuota, legge da standard input
delle righe che contengono ognuna tre interi che rappresentano, nell'ordine,

giorno ora-inizio ora-fine

(come separatore si usano degli spazi, eventualmente ripetuti)

e per ciascuna aggiunge in agenda il corrispondente appuntamento, se valido e non in sovrapposizione con altri.
Una volta raggiunto EOF (che su standard input, su sistemi Linux, si ottiene premendo CTRL D), il programma visualizza su standard output gli appuntamenti in agenda, nell'ordine in cui sono stati inseriti.

Si assuma che da standard input siano inseriti solo numeri interi (non occorre fare controlli).

*/
