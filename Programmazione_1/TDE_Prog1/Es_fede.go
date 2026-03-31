package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var parola string
	var lettere [53]string = [53]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}
	var trova string

	fmt.Println("Inserire una frase: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parola = scanner.Text()
		break
	}

	fmt.Println()

	for _, el := range parola {
		for _, lettera := range lettere {
			fmt.Printf("%s%s \n", trova, lettera)

			if string(el) == lettera {
				trova += lettera
				break
			}
		}
	}
}
