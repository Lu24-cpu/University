/*
Leggere da un file, dove le righe erano titoli di qualcosa.
Nell'output ogni titolo doveva comparire tante volte quante le parole,
e tutte le righe dell'output dovevano essere ordinate secondo ordine alfabetico ma considerando ogni parola del titolo.
Es: "via col vento", compare 3 volte la prima tra le prime perché a "col" le altre 2 tra le ultime perché si considera "via" e "vento"
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parole := strings.Split(scanner.Text(), " ")

		sort.Strings(parole)
		fmt.Println(parole)
		for i, _ := range parole {
			fmt.Print(i)
			for _, parola := range parole {
				fmt.Print(parola, " ")
			}
			fmt.Println()
		}
	}
}
