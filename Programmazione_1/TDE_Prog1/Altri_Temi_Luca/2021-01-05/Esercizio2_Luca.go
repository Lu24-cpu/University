// testo esercizio su pdf

package main

import (
	"fmt"
	"os"
)

type MAPPA = map[string]int

func main() {
	occorrenze := make(MAPPA)
	var sorting []string
	parola := os.Args[1]

	caratteri := []rune(parola)
	for i := 0; i < len(caratteri); i++ {
		for j := i + 2; j < len(caratteri); j++ {
			if caratteri[i] == caratteri[j] {
				occorrenze[parola[i:j+1]]++
				if occorrenze[parola[i:j+1]] <= 1 {
					sorting = append(sorting, parola[i:j+1])
				}
			}
		}
	}

	for i, c := range sorting {
		carattere := []rune(c)
		for j := 0; j < len(sorting); j++ {
			car := []rune(sorting[j])
			if len(carattere) > len(car) {
				sorting[i], sorting[j] = sorting[j], sorting[i]
			}
		}
	}

	for _, sottostringa := range sorting {
		fmt.Println(sottostringa, "-> Occorrenze:", occorrenze[sottostringa])
	}
}
