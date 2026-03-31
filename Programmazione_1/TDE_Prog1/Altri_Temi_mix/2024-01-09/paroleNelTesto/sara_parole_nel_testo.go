package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testo := make(map[string][]int)
	var count int
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		count++
		riga := scanner.Text()
		//for _, parola := range riga {
		testo[riga] = append(testo[riga], count)
		//}
	}
	scanner2 := bufio.NewScanner(os.Stdin)
	for scanner2.Scan() {
		riga := scanner.Text()
		if riga[:2] == "? " {
			fmt.Println("righe: ", testo[riga[2:len(riga)-1]])
			/*for parola := range testo{
			    if riga[2:len(riga)-1] == parola{
			        fmt.Println("righe: ", testo[parola])
			        break
			    }
			}*/
		} else {
			fmt.Println("richiesta non valida: ", riga)
		}

	}
}
