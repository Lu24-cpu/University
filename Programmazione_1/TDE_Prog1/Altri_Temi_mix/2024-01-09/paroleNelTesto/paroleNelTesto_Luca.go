// Testo in fondo all'esercizio
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MAPPA map[string][]int

func Richieste(r, d []string) (richieste MAPPA, err []string) {
	richieste = make(MAPPA)
	for _, s := range d {
		parola := []rune(s)
		if string(parola[0]) == "?" {
			for i, riga := range r {
				if strings.Contains(riga, string(parola[2:])) {
					richieste[string(parola[2:])] = append(richieste[string(parola[2:])], i+1)
				}
			}
		} else {
			sl := "richiesta non valida: " + string(parola)
			err = append(err, sl)
		}
	}
	return
}

func Text(righe []string) (parole MAPPA) {
	parole = make(MAPPA)
	for i, riga := range righe {
		riga = strings.TrimLeft(riga, "\t")
		riga = strings.ReplaceAll(riga, "\t", " ")
		parts := strings.Split(riga, " ")
		for _, parola := range parts {
			if parola != "" {
				parole[parola] = append(parole[parola], i+1)
			}
		}
	}
	return
}

func main() {
	var righe, richiesta []string

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err == nil {
		scanner1 := bufio.NewScanner(file)
		for scanner1.Scan() {
			righe = append(righe, scanner1.Text())
		}

		scanner2 := bufio.NewScanner(os.Stdin)
		for scanner2.Scan() {
			richiesta = append(richiesta, scanner2.Text())
		}

		richieste, err := Richieste(righe, richiesta)

		fmt.Println("Richieste:")
		for _, errore := range err {
			fmt.Println(errore)
		}
		for parola, riga := range richieste {
			fmt.Println("Parola: ", parola)
			fmt.Println("Righe: ", riga)
		}

		fmt.Println("\nTesto:")
		paroletext := Text(righe)

		for parola, riga := range paroletext {
			fmt.Println("Parola: ", parola)
			fmt.Println("Righe: ", riga)
		}

	} else {
		fmt.Println("Errore nell'apertura del file")
	}
}

/*
Parole nel testo
----------------

Scrivere un programma paroleNelTesto.go che legge da file (il cui nome viene passato come parametro su linea di comando) un testo e legge poi da standard input un numero arbitrario di richieste, terminate da EOF (ctrl D su nuova riga), relativamente ai numeri di riga (anche ripetuti) in cui una data parola appare nel testo letto da file.

In particolare da standard input il programma riceve richieste nel formato:

	? <parola da cercare>

Il programma stampa "RICERCHE" e poi, per ogni richiesta, su due righe

	"parola:" seguito da uno spazio e dalla parola fornita
	"righe:" seguito da uno spazio e dalla lista (tra parentesi quadre) dei numeri di riga in cui è comparsa quella parola nel testo letto. Se una data parola appare più volte in una riga, il numero di riga apparirà ripetuto.

Se la richiesta NON inizia con '?' (che è poi seguito da uno spazio), il programma stampa "richiesta non valida" seguito dal testo della richiesta e prosegue ignorandola.
Nota: il programma deve funzionare anche per parole non presenti nel testo, nel qual caso la lista sarà vuota.

Infine il programma stampa "TESTO" e tutte le parole del testo, in ordine lessicografico, una per riga, ciascuna seguita da ':', spazio, e dall’elenco (tra parentesi quadre) dei numeri di riga (anche ripetuti) in cui la parola stessa è comparsa nel testo letto (vedi esempio).

Se il programma viene lanciato con un numero di argomenti diverso da 1,
deve terminare stampando "Fornire un nome di file".
Se invece il file non esiste, il programma deve terminare stampando "File non accessibile".

Esempio di esecuzione
---------------------

$ ./paroleNelTesto  medio.file < medio.input
RICERCHE
parola: ha
righe [1]
richiesta non valida: ! gonna
parola: la
righe [1 1 3]
parola: befana
righe [1 3]
parola: anno
righe []
TESTO
befana: [1 3]
e: [1]
fazzoletto: [1]
gonna: [1]
ha: [1]
il: [1]
la: [1 1 3]
ma: [2]
poverina: [2]
quest'anno: [2]
raffreddata: [3]
rattoppata: [1]
è: [3]
*/
