package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Dizione(parola, nomefile string) (sinonimi []string) {
	dizionario, _ := os.Open(nomefile)

	scannerDiz := bufio.NewScanner(dizionario)
	for scannerDiz.Scan() {
		parti := strings.Split(scannerDiz.Text(), ": ")
		if parti[0] == parola {
			partsin := strings.Split(parti[1], ", ")
			sinonimi = append(sinonimi, partsin...)
			sort.Strings(sinonimi)
			/*fmt.Print(sinonimi)
			  sinonimi = sinonimi[:0]*/
		}
	}
	return
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 file names, please")
		return
	}

	testo, errTes := os.Open(os.Args[2])
	defer testo.Close()
	dizionario, errDiz := os.Open(os.Args[1])

	if errDiz != nil {
		fmt.Println("file 1 not found")
		return
	} else if errTes != nil {
		fmt.Println("file 2 not found")
		return
	}

	dizionario.Close()
	scannerTes := bufio.NewScanner(testo)

	var sinonimi []string
	for scannerTes.Scan() {
		parts := strings.Split(scannerTes.Text(), " ")
		for _, parola := range parts {
			fmt.Print(parola)
			sinonimi = Dizione(parola, os.Args[1])

			if len(sinonimi) != 0 {
				fmt.Print(sinonimi)
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
