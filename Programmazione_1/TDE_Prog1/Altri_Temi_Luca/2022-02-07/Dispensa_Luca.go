// testo a fine codice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MAPPA map[string]float64

func aggiornaDispensa(dispensa MAPPA, nomefile string) bool {
	file, err := os.Open(nomefile)
	defer file.Close()

	if err != nil {
		return true
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		riga := scanner.Text()

		parts := strings.Split(riga, ",")

		if parts[0] == "uso" {
			kg, _ := strconv.ParseFloat(parts[2], 64)
			dispensa[parts[1]] -= kg
			if dispensa[parts[1]] < 0 {
				return false
			}
		} else if parts[0] == "approv" {
			kg, _ := strconv.ParseFloat(parts[2], 64)
			dispensa[parts[1]] += kg
		}
	}
	return false
}

func rimanenza(dispensa MAPPA, alimento string) float64 {
	return dispensa[alimento]
}

func main() {
	nomefile := os.Args[1]
	richieste := os.Args[2:]
	dispensa := make(map[string]float64)

	if richieste == nil {
		fmt.Println("file corretto")
		return
	}

	flag := aggiornaDispensa(dispensa, nomefile)

	if flag {
		fmt.Println("File non leggibile")
	}

	for _, richiesta := range richieste {
		fmt.Println(richiesta, rimanenza(dispensa, richiesta))
	}

}

/*

Dispensa
--------

Si scriva un programma (il file deve chiamarsi 'dispensa.go') per gestire una dispensa di un ristorante.

Il programma leggerà da un file CSV (Comma Separated Values) un "diario" degli acquisti e consumi per poi offrire all'utente la possibilità di conteggiare le rimanenze.

Il programma dovrà essere dotato delle seguenti:

- una funzione AggiornaDispensa(dispensa map[string]float64, filename string) bool
	che aggiorna una dispensa prendendo i dati da un file CSV il cui nome è passato come parametro.
	In particolare il file conterrà gli acquisti nel formato "approv,<prodotto>,<peso>" e i consumi nel formato "uso,<prodotto>,<peso>", dove le quantità (pesi) sono sempre espresse in Kg. (Non vanno fatti controlli sul formato dei dati, si può assumere che siano sempre esattamente come nel file di esempio).
	Se il file non esiste o non è leggibile, la funzione deve terminare subito restituendo 'false'.
	Nel caso in cui un prodotto dovesse "andare in negativo"  durante l'aggiornamento della dispensa, la funzione dovrà terminare subito e restituire 'false', mentre restituirà 'true' in caso contrario.

- una funzione Rimanenza(dispensa map[string]float64, alimento string) float64
	che, data una dispensa opportunamente popolata di provviste, restituisce la rimanenza in Kg dell'alimento passato come parametro; restituisce 0 se il prodotto non è presente in lista.

- una funzione main()
	che legge ed elabora un file in formato CSV che contiene i dati relativi ad acquisti e consumi di una dispensa.
	Il nome del file è passato come primo argomento sulla linea di comando.
	In caso di presenza di argomenti oltre al nome del file, il programma stampa poi la rimanenza degli alimenti passati come argomenti successivi.
	In caso di mancanza di argomenti oltre al nome del file, il programma deve stampare solo "file corretto" se nessun prodotto è andato in negativo, e "file non corretto" in caso contrario.


Esempi di esecuzione
---------------------

$ ./dispensa dispensa.csv zucchero miele peperoncino
zucchero 14.177
miele 28.046
peperoncino 9.452

$ ./dispensa dispensaErrata.csv
file non corretto

$ ./dispensa dispensa.csv
file corretto


*/
