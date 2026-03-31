package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Funzione che estrae le parole che corrispondono all'espressione regolare
func estraiParole(testo, r_exp string) []string {
	var parole []string
	re := regexp.MustCompile(r_exp)
	// Troviamo tutte le parole nel testo
	parole = re.FindAllString(testo, -1)
	return parole
}

func main() {
	var parole []string
	var qt []int
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riga := scanner.Text()
		parole = append(parole, estraiParole(riga, `^[A-Z]([a-z][A-Z])*([a-z]|[A-Z])$`)...)
		qt = append(qt, len(estraiParole(riga, `^[A-Z]([a-z][A-Z])*([a-z]|[A-Z])$`)))
	}

	for k, _ := range qt {
		fmt.Println("Alla riga ", k+1, " ci sono ", qt[k], " parole con il pattern")
	}
	fmt.Println("parole trovate", parole)
}
