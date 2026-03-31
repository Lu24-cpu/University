package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sostituisci(s, old, new []byte, n int) []byte {

	vecchio := string(old)

	slice := []string{}

	var c int
	nuovo := string(new)

	parole := strings.Split(string(s), " ")

	ritorna := []byte{}

	for i := 0; i < len(parole); i++ {

		parola := string(parole[i])

		for j := 0; j < len(parola); j++ {

			for k := j + 1; k < len(parola)+1; k++ {

				sottostringa := string(parola[j:k])

				if sottostringa == vecchio {

					c++
				}
			}

		}

		if c == n {

			slice = append(slice, nuovo)
			c--

		} else {

			slice = append(slice, parola)
		}

	}
	for _, i := range slice {

		bytee := []byte(i)

		bytee = append(bytee, ' ')

		ritorna = append(ritorna, bytee...)
	}

	return ritorna

}

func main() {

	frase1 := []byte(os.Args[1])

	frase := string(frase1)

	old := []byte(os.Args[2])

	new := []byte(os.Args[3])

	n, _ := strconv.Atoi(os.Args[4])

	fmt.Println(frase)

	fmt.Println(string((sostituisci(frase1, old, new, n))))

}
