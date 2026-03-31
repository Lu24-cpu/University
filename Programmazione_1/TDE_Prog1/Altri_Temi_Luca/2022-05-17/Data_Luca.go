// Testo Esercizio a fine codice

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckGiorno(anno, mese, gg string) bool {
	year, _ := strconv.Atoi(anno)
	month, _ := strconv.Atoi(mese)
	day, _ := strconv.Atoi(gg)
	if year < 1900 || year > 2100 || month < 1 || month > 12 {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if day < 1 || day > 31 {
				return false
			}
			return true
		case 2:
			if ((day < 1 || day > 29) && ((year%4 == 0 && year%100 != 0) || year%400 == 0)) || (day < 1 || day > 28) {
				return false
			}
			return true
		default:
			if day < 1 || day > 30 {
				return false
			}
			return true
		}

		return false
	}
	return true
}

func CheckOra(ora, minuti, secondi string) bool {
	hour, _ := strconv.Atoi(ora)
	minutes, _ := strconv.Atoi(minuti)
	seconds, _ := strconv.Atoi(secondi)

	if hour < 0 || hour > 23 || minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		return false
	}
	return true
}

func main() {
	mesi := [12]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"}
	data := os.Args[1]

	parts := strings.Split(data, " ")
	giorno, ora := parts[0], parts[1]

	giornop := strings.Split(giorno, "-")
	orap := strings.Split(ora, ":")

	flag1 := CheckGiorno(giornop[0], giornop[1], giornop[2])
	flag2 := CheckOra(orap[0], orap[1], orap[2])
	if flag1 && flag2 {
		mese, _ := strconv.Atoi(giornop[1])
		fmt.Println(giornop[2], mesi[mese-1], giornop[0])
	}
}

/**

Esercizio "data"
================

Scrivere un programma Go (il file deve chiamarsi 'data.go') che:

- legge da linea di comando una stringa del tipo

	"2019-06-05 14:11:25"

  NB: anno-mese-giorno ore:minuti:secondi, *in un solo argomento tra virgolette*

- determina se i formati della data e dell'ora sono validi:
	- mese tra 01 e 12, giorno tra 01 e 31, anno tra 1900 e 2100, con separatore '-'
	  e giorno e mese di due caratteri
	- ore, minuti e secondi di due caratteri, con separatore ':'

- determina se l'ora è valida (ora tra 00 e 23, minuti e secondi tra 00 e 59)

- se i formati della data e dell'ora sono validi e l'ora è valida, stampa la data nel formato

	gg mese aaaa

con il nome del mese in italiano; altrimenti stampa il messaggio "argomento non valido"


Esempi
------

$ go run data.go "2019-06-05 12:11:25"
05 giugno 2019

$ go run data.go "2019-06-55 12:11:25"
argomento non valido

$ go run data.go "2019-06-05 12-11-25"
argomento non valido

$ go run data.go "2019-11-30 12:11:25"
30 novembre 2019

$ go run data.go "2019-11-30 12:11:75"
argomento non valido

$ go run data.go "2019-13-31 12:11:25"
argomento non valido

*/
