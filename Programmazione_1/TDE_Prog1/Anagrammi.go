package main

import (
	"fmt"
	"os"
)

func Anagrammi(risposta []string, parola, s string) []string {
	t := []rune(s)
	if len(t) == 0 {
		risposta = append(risposta, parola)
		return risposta
	}
	for i := 0; i < len(t); i++ {
		risposta = Anagrammi(risposta, parola+string(t[i]), string(s[i+1:]))
	}
	return risposta
}

func RimuoviDop(risposte []string) (reduced1 []string) {
	ridotto := make(map[string]bool)

	for _, word := range risposte {
		if !ridotto[word] {
			ridotto[word] = true
		}
	}

	for key, _ := range ridotto {
		if len(key) == 6 {
			reduced1 = append(reduced1, key)
		}
	}

	return
}

func main() {
	s := os.Args[1]

	ris := Anagrammi(nil, "", s)

	fmt.Println(RimuoviDop(ris))
}
