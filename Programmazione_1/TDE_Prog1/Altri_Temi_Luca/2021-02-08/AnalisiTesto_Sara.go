package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Vocabolario map[string]int

func (voc Vocabolario) String() {
	var sl []string
	for chiavi := range voc {
		sl = append(sl, chiavi)
	}
	sort.Strings(sl)
	for _, parole := range sl {
		fmt.Println(parole, " : ", voc[parole])
	}
}

func statistiche(voc Vocabolario) (int, int, float64) {
	var basso, alto, somma, i int
	var media float64
	for _, val := range voc {
		if i == 0 {
			basso, alto = val, val
			i++
		}
		if val < basso {
			basso = val
		}
		if val > alto {
			alto = val
		}
		somma += val
	}
	media = float64(somma) / float64(len(voc))
	return basso, alto, media
}

func selezione(voc Vocabolario, char rune) []string {
	var sl []string
	for chiave := range voc {
		var sr []rune
		for _, carattere := range chiave {
			sr = append(sr, carattere)
		}
		if len(sr) != 0 {
			if sr[0] == char {
				sl = append(sl, chiave)
			}
		}
	}
	return sl
}

func main() {
	voc := make(Vocabolario)
	if len(os.Args) < 3 {
		fmt.Println("indicare il nome del file e iniziale")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	c := os.Args[2]
	var char rune
	for _, carattere := range c {
		char = carattere
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			continue
		}
		riga = strings.TrimSpace(riga)
		parts := strings.Split(riga, " ")
		for _, parola := range parts {
			if parola != "" {
				voc[parola]++
			}
		}
	}
	voc.String()
	fmt.Println(statistiche(voc))
	fmt.Println(selezione(voc, char))
}
