package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operazione struct {
	data    string
	tipo    string
	importo int
}

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {

		fmt.Println("errore nell'apertuta del file")
		return
	}

	mappa := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		riga := strings.Split(scanner.Text(), ";")

		data := strings.Split(riga[0], "_")

		mese, _ := strconv.Atoi(data[1])

		giorno, _ := strconv.Atoi(data[2])

		datareale := fmt.Sprintf("%s_%d_%d", data[0], mese, giorno)

		importo, _ := strconv.Atoi(riga[2])

		operazione := Operazione{datareale, riga[1], importo}

		if operazione.tipo == "V" {

			mappa[operazione.data] += operazione.importo
		} else if operazione.tipo == "P" {

			mappa[operazione.data] -= operazione.importo
		}

	}

	for i, v := range mappa {

		riga := strings.Split(i, "_")

		fmt.Printf("%s/%s/%s -> %d\n", riga[0], riga[1], riga[2], v)
	}
}
