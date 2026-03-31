// testo esercizio nel pdf

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func generateSubsequences(s string, r int, current string, result *[]string) {
	if r == 0 {
		*result = append(*result, current+s)
		return
	}
	if len(s) == 0 {
		return
	}
	// Escludi il primo carattere
	generateSubsequences(s[1:], r-1, current, result)
	// Includi il primo carattere
	generateSubsequences(s[1:], r, current+string(s[0]), result)
}

func sottosequenze(s string, n int) []string {
	var stsequenza []string
	lunghezza := len(s)
	for i := 0; i <= lunghezza-1; i++ {
		for k := i + 1; k <= lunghezza; k++ {
			if k > lunghezza {
				continue
			}
			t := n - ((i - 1) + (k - i))
			if -t > 0 {
				if -t < i {
					parola := s[:-t] + string(s[i]) + s[k:]
					stsequenza = append(stsequenza, parola)
				}
				continue
			}
			parola := string(s[i]) + s[k:len(s)-t]
			stsequenza = append(stsequenza, parola)
		}
	}
	return stsequenza
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <stringa> <numero>")
		return
	}

	num := os.Args[1]
	r, _ := strconv.Atoi(os.Args[2])

	if r <= 0 || r >= len(num) {
		fmt.Println("Il numero deve essere maggiore di 0 e minore della lunghezza della stringa.")
		return
	}

	// Genera tutte le sottosequenze rimuovendo esattamente r caratteri
	var sottoseq []string
	generateSubsequences(num, r, "", &sottoseq)

	sort.Strings(sottoseq)
	fmt.Println(sottoseq[0])
	sottoseq = sottosequenze(num, r)
	sort.Strings(sottoseq)
	fmt.Println(sottoseq[0])
}
