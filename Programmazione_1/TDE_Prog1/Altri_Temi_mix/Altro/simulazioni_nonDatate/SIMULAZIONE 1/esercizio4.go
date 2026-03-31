package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

//Funzione che data una slice di stringe ritorna tutte le combinazioni di slice
//di stringhe con le rimozioni
func rimuoviCifra(nList []string) []string {
	var res []string
	for _, n := range nList {
		for i := range n {
			res = append(res, n[:i]+n[i+1:])
		}
	}
	return res
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Troppi pochi argomenti")
		os.Exit(1)
	}
	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])

	//RIPETO PER D VOLTE RIMUOVENDO OGNI VOLTA UNA CIFRA
	res := []string{n}
	for i := 0; i < d; i++ {
		res = rimuoviCifra(res)

	}
	//TROVO IL MINIMO
	min := math.MaxInt32
	for _, i := range res {
		k, _ := strconv.Atoi(i)
		if k < min {
			min = k
		}
	}
	fmt.Println(min)
}
