package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numero, _ := strconv.Atoi(os.Args[1])

	cento := numero / 100

	numero = numero - (100 * cento)

	fmt.Printf("Taglio : %s Quantità: %d\n", "100", cento)

	cinquanta := numero / 50

	numero = numero - (50 * cinquanta)

	fmt.Printf("Taglio : %s Quantità: %d\n", "50", cinquanta)

	venti := numero / 20

	numero = numero - (20 * venti)

	fmt.Printf("Taglio : %s Quantità: %d\n", "20", venti)

	dieci := numero / 10

	numero = numero - (10 * dieci)

	fmt.Printf("Taglio : %s Quantità: %d\n", "10", dieci)

	cinque := numero / 5

	numero = numero - (5 * cinque)

	fmt.Printf("Taglio : %s Quantità: %d\n", "5", cinque)

	due := numero / 2

	numero = numero - (2 * due)

	fmt.Printf("Taglio : %s Quantità: %d\n", "2", due)

	uno := numero / 1

	numero = numero - (1 * uno)

	fmt.Printf("Taglio : %s Quantità: %d\n", "1", uno)

}
