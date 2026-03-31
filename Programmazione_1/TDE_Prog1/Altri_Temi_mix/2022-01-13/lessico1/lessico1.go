package main

import (
	"fmt"
	"bufio"
	"os"
)

const MEMORIZE = "+"
const CONSULT = "?"
const MINMAX = "m"
const PRINT = "p"

func main(){
	var action string

	fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		riga := scanner.Text()
		action = string(riga[0])
		
		switch action{
		case MEMORIZE:
			fmt.Println("alimento il dizionario")
		case CONSULT:
			fmt.Println("consulto il dizionario")
		case MINMAX:
			fmt.Println("lunghezza min e max")
		case PRINT:
			fmt.Println("stampo il dizionario ordinato")
		default:
			fmt.Println("opzione non valida")
		}
	}
}