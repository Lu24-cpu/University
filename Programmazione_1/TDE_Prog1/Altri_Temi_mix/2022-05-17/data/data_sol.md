package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	diviso := strings.Split(os.Args[1], " ")

	data := strings.Split(diviso[0], "-")
	ora := strings.Split(diviso[1], ":")

	var mese string

	if data[0] >= "1900" && data[0] <= "2100" {

		if data[1] >= "01" && data[1] <= "12" && len(data[1]) == 2 {

			if data[2] >= "01" && data[2] <= "31" && len(data[2]) == 2 {

				if string(os.Args[1][4]) == "-" && string(os.Args[1][7]) == "-" {

					if ora[0] >= "00" && ora[0] <= "23" && len(ora[0]) == 2 {

						if ora[1] >= "00" && ora[1] <= "59" && len(ora[1]) == 2 {

							if ora[2] >= "00" && ora[2] <= "59" && len(ora[2]) == 2 {

								if string(os.Args[1][13]) == ":" && string(os.Args[1][16]) == ":" {

									switch data[1] {

									case "01":
										mese = "gennaio"

									case "02":
										mese = "febbraio"
									case "03":
										mese = "marzo"
									case "04":
										mese = "aprile"
									case "05":
										mese = "maggio"
									case "06":
										mese = "giugno"
									case "07":
										mese = "luglio"
									case "08":
										mese = "agosto"
									case "09":
										mese = "settembre"
									case "10":
										mese = "ottobre"
									case "11":
										mese = "novembre"
									case "12":
										mese = "dicembre"

									}

									fmt.Println(data[2], mese, data[0])
									return

								}

							}

						}

					}

				}
			}
		}

	}

	fmt.Println("argomento non valido")

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
