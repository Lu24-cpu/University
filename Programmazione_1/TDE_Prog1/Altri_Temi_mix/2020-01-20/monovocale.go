package main

import "fmt"

func main() {
	var s string

	for {

		var c int

		fmt.Println("Inserisci una parola")

		fmt.Scan(&s)

		mappa := make(map[rune]int)

		if s == "0" {

			break
		}

		for _, i := range s {

			if i < 'a' || i > 'z' {

				fmt.Println("errore")
				c++

				break
			}

			switch i {

			case 'a':
				mappa['a']++

			case 'e':
				mappa['e']++
			case 'i':
				mappa['i']++
			case 'o':
				mappa['o']++
			case 'u':
				mappa['u']++

			}
		}

		if c == 0 {

			if len(mappa) == 1 {

				for chiave, valore := range mappa {

					fmt.Println("corretto", string(chiave), valore)
				}
			} else {

				fmt.Println("sbagliato")

			}
		} else {

			c--
		}

	}
}
