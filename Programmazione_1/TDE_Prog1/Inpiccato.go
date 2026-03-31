package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ParolaRandom() string {
	parole := []string{"parato", "prato", "aurato", "laureato", "parate", "tarpea", "paura", "para", "apra", "telatura", "aurate", "liftare", "laureate", "reato", "laureo", "areate", "pare", "pera", "rape", "ratto", "ratti", "urlerà", "trailer", "urlerai", "feto", "rurali", "aurei", "rara", "urli", "aree"}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(30)

	return parole[n]
}

func TrovaParola(word string) {
	var countlett int
	var lettera rune
	parola := []rune(word)

	fmt.Println("La parola è lunga", len(parola))

	for i := 1; i <= 10; i++ {
		var flag bool = false

		fmt.Print("Inserire una lettera:")
		fmt.Scan(&lettera)

		for _, elemento := range parola {
			if lettera != elemento {
				fmt.Printf("_ ", lettera)
			} else {
				fmt.Printf("%c ", lettera)
				countlett++
				flag = true
			}
		}
		fmt.Println()
		if !flag {
			fmt.Println("Lettera inserita non è presente")
		}

	}
}

func main() {
	fmt.Println("Benvenuto/i al gioco dell'inpiccato")

	TrovaParola(ParolaRandom())

}
