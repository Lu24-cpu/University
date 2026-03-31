package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
)

func LeggiSinonimi(file *os.File) (mappaSinonimi map[string][]string){
	mappaSinonimi = make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		riga := scanner.Text()
		indiceSeparatore := strings.Index(riga, SEP_1)
		word := riga[:indiceSeparatore]
		listaSinonimi := strings.Split(riga[indiceSeparatore+2:], SEP_2)	// resto della riga senza contare primo spazio
		sort.Strings(listaSinonimi)
		mappaSinonimi[word] = listaSinonimi
	}
	return
}

func ElaboraTesto(file *os.File, mappaSinonimi map[string][]string){
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		listaParole := strings.Fields(scanner.Text())
		for _, parola := range listaParole{
			fmt.Print(parola)
			sinonimi, ok := mappaSinonimi[parola]
			if ok{
				fmt.Print(sinonimi)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

const SEP_1 = ": "
const SEP_2 = ", "

func main(){
	if len(os.Args) != 3{
		fmt.Println("2 file names, please")
		return
	}
	fileSinomini, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("file 1 not found")
		return
	}
	defer fileSinomini.Close()

	mappaSinonimi := make(map[string][]string)
	mappaSinonimi = LeggiSinonimi(fileSinomini)

	fileTesto, err := os.Open(os.Args[2])
	if err != nil{
		fmt.Println("file 2 not found")
		return
	}
	defer fileTesto.Close()

	ElaboraTesto(fileTesto, mappaSinonimi)

}
