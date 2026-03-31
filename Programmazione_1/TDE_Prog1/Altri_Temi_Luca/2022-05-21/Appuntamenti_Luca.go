// Testo esercizio a fine Codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type APPUNTAMENTO struct {
	giorno, oraInizio, oraFine int
}

func NuovoAppuntamento(gg, h1, h2 int) (APPUNTAMENTO, bool) {
	if gg >= 1 && gg <= 366 && h1 >= 0 && h1 <= h2 && h2 <= 24 {
		x := APPUNTAMENTO{gg, h1, h2}
		return x, true
	}
	x := APPUNTAMENTO{0, 0, 0}
	return x, false
}

func CheckSovrapposizioneAppuntamenti(app1, app2 APPUNTAMENTO) bool {
	if app1.giorno != app2.giorno {
		return false
	}
	return !(app1.oraFine <= app2.oraInizio || app2.oraFine <= app1.oraInizio)
}

func AggiungiAppuntamento(app APPUNTAMENTO, agenda *[]APPUNTAMENTO) bool {
	for _, app1 := range *agenda {
		if CheckSovrapposizioneAppuntamenti(app1, app) {
			return false
		}
	}

	*agenda = append(*agenda, app)
	return true
}

func AppuntamentiGiornata(gg int, agenda []APPUNTAMENTO) (appgg []APPUNTAMENTO) {
	for _, app := range agenda {
		if gg == app.giorno {
			appgg = append(appgg, app)
		}
	}
	return
}

func main() {
	var agenda []APPUNTAMENTO
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := scanner.Text()
		segno := strings.Fields(riga)

		gg, _ := strconv.Atoi(segno[0])
		h1, _ := strconv.Atoi(segno[1])
		h2, _ := strconv.Atoi(segno[2])

		appuntamento, flag := NuovoAppuntamento(gg, h1, h2)
		if flag {
			AggiungiAppuntamento(appuntamento, &agenda)
		}
	}
	for _, appuntamento := range agenda {
		fmt.Print(appuntamento, "\n")
	}
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
