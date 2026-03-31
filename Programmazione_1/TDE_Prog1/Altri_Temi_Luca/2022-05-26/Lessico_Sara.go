package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("+ (Legge una riga e memorizza le parole)\n? (Legge una parola e indica le righe che la contengono)\np (Stampa le parole)")
	dizionario := make(map[string][]int)
	var rigo int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		switch riga[0] {
		case '+':
			rigo++
			parts := strings.Split(riga[1:], " ")
			for _, parola := range parts {
				dizionario[parola] = append(dizionario[parola], rigo)
			}
		case '?':
			fmt.Printf("parola: %v\nrighe %v\n", riga[1:], dizionario[riga[1:]])
		case 'p':
			fmt.Println(dizionario)
		}
	}
}
