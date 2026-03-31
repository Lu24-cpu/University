package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
)

func AddToDict(line string, n int, dict map[string][]int){
	line = line[2:]
	listaParole := strings.Fields(line)
	for _, parola := range listaParole{
		dict[parola] = append(dict[parola], n)
	}
}

func Stats(dict map[string][]int) (int, int){
	// lunghezza min e max delle parole memorizzate
	var min, max int

	for parola, _ := range dict{
		if len(parola) >= max{
			max = len(parola)
		}
		if len(parola)<=min || min == 0 {
			min = len(parola)
		}
	}
	return min, max
}

func PrintMenu(){
	fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")
}
func PrintDict(m map[string][]int){
	var keySlice []string
	for key, _ := range m{
		keySlice = append(keySlice, key)
	}
	sort.Strings(keySlice)
	for _, key := range keySlice{
		fmt.Println(key, ":", m[key])
	}
}

func ConsultDict(line string, dict map[string][]int){
	line = line[2:]
	fmt.Println("parola:", line)
	fmt.Println("righe", dict[line])
}

func PrintMinMax(dict map[string][]int){
	min, max := Stats(dict)
	fmt.Println(min, max)
}

const MEMORIZE = "+"
const CONSULT = "?"
const MINMAX = "m"
const PRINT = "p"

func main(){
	var currRiga int
	var dizionario map[string][]int
	dizionario = make(map[string][]int)

	currRiga = 1
	PrintMenu()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		riga := scanner.Text()

		switch string(riga[0]){
		case MEMORIZE:
			fmt.Println("alimento il dizionario")
			AddToDict(riga, currRiga, dizionario)
			currRiga++
		case CONSULT:
			fmt.Println("consulto il dizionario")
			ConsultDict(riga, dizionario)
		case MINMAX:
			fmt.Println("lunghezza min e max")
			PrintMinMax(dizionario)
		case PRINT:
			fmt.Println("stampo il dizionario ordinato")
			PrintDict(dizionario)
		default:
			fmt.Println("opzione non valida")
		}
	}
}