package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// EstraiParole estrae tutte le parole che corrispondono alla regex r_exp da un testo.
func EstraiParole(testo, r_exp string) (parole []string) {
	re := regexp.MustCompile(r_exp)
	parole = re.FindAllString(testo, -1)
	return
}

func main() {
	// Nome del file passato come argomento da linea di comando
	nome := os.Args[1]
	var parole []string
	var quantita []int

	// La regex per trovare l'alternanza maiuscola-minuscola
	r_exp := `^[A-Z][a-z]*([A-Z][a-z]*)*$`

	// Apertura del file
	file, err := os.Open(nome)
	if err != nil {
		fmt.Println("Errore nell'aprire il file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Lettura e analisi riga per riga
	for scanner.Scan() {
		riga := scanner.Text()

		// Estrazione delle parole che corrispondono alla regex
		paroleRiga := EstraiParole(riga, r_exp)
		parole = append(parole, paroleRiga...)

		// Calcolo quante parole nella riga corrispondono al pattern
		parts := strings.Fields(riga)
		var count int
		for _, word := range parts {
			if len(EstraiParole(word, r_exp)) > 0 {
				count++
			}
		}
		quantita = append(quantita, count)
	}

	// Gestione degli errori di scansione
	if err := scanner.Err(); err != nil {
		fmt.Println("Errore nella lettura del file:", err)
	}

	// Stampa i risultati
	for i, count := range quantita {
		fmt.Printf("riga %d : %d parola/e con il pattern\n", i+1, count)
	}

	// Stampa tutte le parole trovate
	fmt.Println("stringhe trovate :", parole)
}
