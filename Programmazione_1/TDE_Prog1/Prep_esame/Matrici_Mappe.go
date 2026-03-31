// crea un esercizio che contenga un metodo che stampi una matrici riempite da Stdin e riunite in una mappa il main deve essere vuoto
// nome, elemento elemento elemento
// divisione con split e trimspace
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MAPPA = map[string][][]int

func Menu() {
	fmt.Println("Inserire:\n1. Per aggiungere una matrice\n2. Per stampare le matrici\n0. Per terminare")
}

func CreaRiga(s string) (riga []int) {
	elementi := strings.Split(s, " ")

	for _, nst := range elementi {
		nint, _ := strconv.Atoi(nst)
		riga = append(riga, nint)
	}

	return
}

func CreaMap() (nome string, mb [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ",") {
			nome = scanner.Text()
			nome = nome[:len(nome)-1]
		} else {
			mb = append(mb, CreaRiga(scanner.Text()))
		}
	}

	return
}

func Stampa(m MAPPA) {
	for key, el := range m {
		fmt.Println(key)
		for _, riga := range el {
			for _, elementi := range riga {
				fmt.Printf("%d\t", elementi)
			}
			fmt.Println()
		}
	}
}

func Input(matrm MAPPA) (MAPPA, bool) {
	var selezione int
	selezione = 1

	fmt.Scan(&selezione)
	switch selezione {
	case 1:
		nome, matrice := CreaMap()
		matrm[nome] = matrice
	case 2:
		Stampa(matrm)
	case 0:
		return matrm, false
	default:
		fmt.Println("Selezione non valida reinserire")
	}
	return matrm, true
}

func main() {
	var flag bool = true
	matrm := make(MAPPA)
	Menu()
	for flag {
		matrm, flag = Input(matrm)
	}
}
