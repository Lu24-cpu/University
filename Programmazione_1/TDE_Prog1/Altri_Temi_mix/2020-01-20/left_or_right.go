package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var dir, dirc string

	var v, d, volte float64

	for {
		fmt.Println("Dove sto guardando?")
		fmt.Scan(&dir)

		if dir == "0" {

			break
		}

		volte++

		n := rand.Intn(100)

		if n%2 == 0 {

			dirc = "d"
		} else {

			dirc = "s"
		}

		if dirc == dir {

			fmt.Println("Giusto, hai vinto")

			v++
		} else {

			fmt.Println("Sbagliato, hai perso!")
			d++
		}

	}

	fmt.Print("Vinto: ", (v/volte)*100.0, "% Perso: ", (d/volte)*100.0, "%\n")

}
