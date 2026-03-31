package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cancella(lista []string) []string {
	var nuova_lista []string
	var da_saltare int
	for _, val := range lista {
		if da_saltare == 0 {
			cifra, err := strconv.Atoi(val)
			if err != nil {
				nuova_lista = append(nuova_lista, val)
				continue // (non è una cifra)
			}
			da_saltare = cifra
		} else {
			da_saltare--
		}
	}

	return nuova_lista
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		return
	}
	fname := os.Args[1]
	fp, err := os.Open(fname)
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}

	var parole []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		riga := strings.Fields(scanner.Text())
		parole = append(parole, riga...)
	}
	fmt.Println(parole)
	fmt.Println(cancella(parole))
}
