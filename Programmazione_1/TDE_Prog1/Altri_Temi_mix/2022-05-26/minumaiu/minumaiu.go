package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func estraiParole(testo, r_exp string) (parole []string) {

	slice := []string{}

	parola := strings.Split(testo, " ")

	for _, parolaa := range parola {

		for i := 0; i < len(parolaa); i++ {

			for t := i + 1; t < len(parolaa); t++ {

				sottostringa := parolaa[i:t]

				if sottostringa == r_exp {

					slice = append(slice, sottostringa)
				}
			}

		}

	}

	return slice

}

func main() {

	file, err := os.Open(os.Args[1])

	slice := []string{}

	if err != nil {

		fmt.Println("inserire nome file valido")
		return
	}

	scanner := bufio.NewScanner(file)

	var sottostringa string

	var line, count int

	for scanner.Scan() {

		line++

		riga := strings.Split(scanner.Text(), " ")

		for _, parole := range riga {

			if parole != "" {

				for i := 0; i < len(parole)-1; i++ {

					if unicode.IsUpper(rune(parole[i])) && unicode.IsUpper(rune(parole[i+1])) == false {

						sottostringa += string(parole[i]) + string(parole[i+1])
						i++

					} else {

						if sottostringa != "" {

							slice = append(slice, sottostringa)
							count++
							sottostringa = ""

						}
					}

				}

				if sottostringa != "" {

					slice = append(slice, sottostringa)
					count++
				}

				sottostringa = ""

			}

		}
		fmt.Printf("riga %d : %d parola/e con il pattern\n", line, count)

		count = 0

	}

	fmt.Println(slice)
}

/*
MinuMaiu
--------

Scrivere un programma minumaiu.go, dotato di:
-  una funzione
	estraiParole(testo, r_exp string) (parole []string)
che estrae da testo le (sotto)stringhe che corrispondono a r_exp
- una funzione
	main()
che legge da un file, il cui nome è passato su linea di comando,
e, per ogni riga del file, stampa quante stringhe contengono
pattern <alternanza di mauiscola-minuscola>.
Alla fine stampa tutte le parole trovate.

Per chiarire, la stringa AbBc e` tutta <alternanza di mauiscola-minuscola>,
ma anche la stringa eSaMelab contiene il pattern
<alternanza di mauiscola-minuscola>.

Esempio
-------

Input (da file):
AbBc

 	casa mia
Ah MmMmMm
oHoH abcdEfGhilmnOp

Output:
riga 1 : 1 parola/e con il pattern
riga 2 : 0 parola/e con il pattern
riga 3 : 0 parola/e con il pattern
riga 4 : 2 parola/e con il pattern
riga 5 : 2 parola/e con il pattern
stringhe trovate : [AbBc Ah MmMmMm Ho EfGh Op]
*/
