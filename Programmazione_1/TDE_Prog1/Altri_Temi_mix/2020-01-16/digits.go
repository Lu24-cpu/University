package main

import "fmt"

func main() {
	var s string

	for {
		fmt.Println("	Inserisci un numero:")
		fmt.Scan(&s)
		mappa := make(map[rune]int)
		var c int
		if s == "0" {
			fmt.Println("	Fine")
			break
		}
		for _, i := range s {
			mappa[i]++
		}
		for _, valore := range mappa {

			if valore > 1 {

				fmt.Println("	Il numero ha cifre ripetute")
				c++

			}
		}
		if c == 0 {
			fmt.Println("	Non ha cifre ripetute")
		}

	}
}
