package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func analyze_text(fname string) map[string][]int {
	fp, err := os.Open(fname)
	if err != nil {
		fmt.Println("File non accessibile")
		os.Exit(1)
	}
	defer fp.Close()

	parole := make(map[string][]int)

	scanner := bufio.NewScanner(fp)
	index_riga := 1 // non è specificato da che valore cominciare a contare righe (da input esempio era 1)
	for scanner.Scan() {
		riga := strings.Fields(scanner.Text())
		for _, parola := range riga {
			parole[parola] = append(parole[parola], index_riga)
		}
		index_riga++
	}
	return parole
}

const SIMBOLO_RICHIESTA = "?"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire un nome di file")
		return
	}

	var parole map[string][]int
	var to_search string
	var simbolo string

	parole = make(map[string][]int)
	// leggi file e salva posizione di riga di ogni parola
	parole = analyze_text(os.Args[1])

	// SEZIONE RICERCHE
	fmt.Println("RICERCHE")
	_, err := fmt.Scan(&simbolo)
	fmt.Scan(&to_search)
	for err != io.EOF {
		if simbolo != SIMBOLO_RICHIESTA {
			fmt.Println("richiesta non valida:", simbolo, to_search)
		} else {
			fmt.Println("parola:", to_search)
			fmt.Println("righe", parole[to_search])
		}

		_, err = fmt.Scan(&simbolo)
		fmt.Scan(&to_search)
	}

	solo_chiavi := []string{}
	for parola, _ := range parole {
		solo_chiavi = append(solo_chiavi, parola)
	}
	sort.Strings(solo_chiavi)
	// SEZIONE TESTO
	fmt.Println("TESTO")
	for _, parola := range solo_chiavi {
		fmt.Printf("%s: ", parola)
		fmt.Println(parole[parola])
	}
}
