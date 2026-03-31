package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var Nomi = make(map[string]string)

type data struct {
	giorno int
	mese   int
	anno   int
}

func LetturaFile() {
	file, err := os.Open("Biblioteca.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Errore, file non trovato, reinserire")
		return
	} else {
		libri := make([]byte, 1024)

		for {
			n, err2 := file.Read(libri)
			if err2 != nil {
				if err2 == io.EOF {
					break
				}
				fmt.Println("Errore lettura file: ", err2)
				return
			}
			fmt.Println(string(libri[:n]))
		}
	}
}

func RitiroLibro() {
	var libreria []string
	var nome, persona string
	var day data

	fmt.Println("Inserire il nome del libro (premere invio e ctrl+d per terminare): ")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan() {
		nome += text.Text()
	}

	fmt.Println("Inserire la data di oggi:")
	fmt.Scan(&day.giorno, &day.mese, &day.anno)

	file, err := os.Open("Biblioteca.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file: ", err)
		return
	}
	defer file.Close()

	filescanner := bufio.NewScanner(file)
	for filescanner.Scan() {
		libreria = append(libreria, filescanner.Text())
	}

	var updatedLibreria []string
	var found bool

	for _, libro := range libreria {
		if strings.Contains(libro, nome) {
			parts := strings.Split(libro, " ")
			numero, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				fmt.Println("Errore nella richiesta del libro", err)
				return
			}

			if numero == 0 {
				fmt.Println("Non è presente il libro cercato")
				return
			}

			numero--
			parts[len(parts)-1] = strconv.Itoa(numero)
			libro = strings.Join(parts, " ")
			found = true
		}
		updatedLibreria = append(updatedLibreria, libro)
	}

	if !found {
		fmt.Println("Libro non trovato")
		fmt.Println("")
		return
	}

	// Riscrive tutto il contenuto aggiornato nel file
	file, err = os.Create("Biblioteca.txt")
	if err != nil {
		fmt.Println("Errore nell apertura del file per la scrittura: ", err)
		fmt.Println("")
		return
	}
	defer file.Close()

	fmt.Print("Inserire il proprio nome: ")
	fmt.Scan(&persona)

	Nomi[persona] = nome

	// Salva i nomi subito dopo l'aggiornamento della mappa
	SalvataggioNomi(day)

	writer := bufio.NewWriter(file)
	for _, libro := range updatedLibreria {
		fmt.Fprintln(writer, libro)
	}
	writer.Flush()

	fmt.Println("Libro ritirato con successo!")
	fmt.Println("")
}

func RestituzioneLibro() {
	var libreria, line []string
	var nome, nomelibro string
	var day, restday data
	var delta int

	fmt.Println("Inserire il proprio nome:")
	fmt.Scanln(&nome)

	fmt.Println("Inserire il nome del libro da restituire (per terminare invio e ctrl+D):")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan() {
		nomelibro += text.Text()
	}

	fmt.Println("Inserire la data di oggi:")
	fmt.Scan(&day.giorno, &day.mese, &day.anno)

	// Leggi il file Prestiti.txt e popola la mappa Nomi
	filePrestiti, err := os.Open("Prestiti.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file Prestiti.txt: ", err)
		return
	}
	defer filePrestiti.Close()

	scannerPrestiti := bufio.NewScanner(filePrestiti)
	for scannerPrestiti.Scan() {
		line = append(line, scannerPrestiti.Text())
	}

	for _, prestiti := range line {
		if strings.Contains(prestiti, nome) && strings.Contains(prestiti, nomelibro) {
			for i := range len(prestiti) {
				if unicode.IsSpace(rune(prestiti[i])) && unicode.IsNumber(rune(prestiti[i+1])) {
					restday.giorno, _ = strconv.Atoi(string(prestiti[i+1]))
					restday.mese, _ = strconv.Atoi(string(prestiti[i+3]))
					restday.anno, _ = strconv.Atoi(string(prestiti[i+5:]))
				}
			}
		}
	}

	// Calcolo del numero di giorni di ritardo
	delta = (restday.anno-day.anno)*365 + (restday.mese-day.mese)*30 + (restday.giorno - day.giorno)
	if delta > 30 {
		fmt.Println("Restituzione libro in ritardo, pagare il periodo in più")
	}

	// Leggi il file Biblioteca.txt e aggiorna la lista dei libri
	file, err := os.Open("Biblioteca.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file Biblioteca.txt: ", err)
		return
	}
	defer file.Close()

	filescanner := bufio.NewScanner(file)
	for filescanner.Scan() {
		libreria = append(libreria, filescanner.Text())
	}

	var updatedLibreria []string
	var found bool = false

	for _, item := range libreria {
		var titoloLibro, numerostr string

		for i := range item {
			if string(item[i]) == "," && unicode.IsSpace(rune(item[i+1])) {
				titoloLibro, numerostr = item[:i], item[i+2:]
				break
			}
		}

		numero, _ := strconv.Atoi(numerostr)

		if titoloLibro == nomelibro {
			numero++
			found = true
		}

		updatedLibreria = append(updatedLibreria, fmt.Sprintf("%s, %d", titoloLibro, numero))
	}

	if !found {
		fmt.Println("Libro non trovato")
		return
	}

	// Riscrivi tutto il contenuto aggiornato nel file Biblioteca.txt
	file, err = os.Create("Biblioteca.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file per la scrittura: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, item := range updatedLibreria {
		fmt.Fprintln(writer, item)
	}
	writer.Flush()

	// Aggiorna il file Prestiti.txt rimuovendo il prestito restituito
	delete(Nomi, nome)

	// Salva i nomi subito dopo l'aggiornamento della mappa
	SalvataggioNomi(day)

	filePrestiti, err = os.Create("Prestiti.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file Prestiti.txt per la scrittura: ", err)
		return
	}
	defer filePrestiti.Close()

	writerPrestiti := bufio.NewWriter(filePrestiti)
	for k, v := range Nomi {
		fmt.Fprintln(writerPrestiti, k, v)
	}
	writerPrestiti.Flush()

	fmt.Println("Libro restituito con successo!")
	fmt.Println("")
}

func MenuBiblioteca() {
	fmt.Println("Inserire:")
	fmt.Println("1. Per vedere i libri disponibili")
	fmt.Println("2. Per prendere in prestito un libro")
	fmt.Println("3. Per restituire un libro")
	fmt.Println("0. Per uscire dal programma")
	fmt.Println("")
}

func MenuInizio() {
	fmt.Println("Inserire:")
	fmt.Println("1. Per vedere i libri disponibili")
	fmt.Println("2. Per prendere in prestito un libro")
	fmt.Println("3. Per restituire un libro")
	fmt.Println("")
}

func SalvataggioNomi(day data) {
	file, err := os.Create("Prestiti.txt")

	if err != nil {
		fmt.Println("Errore nell'apertura del file Prestiti.txt: ", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for nome, prestito := range Nomi {
		fmt.Fprintf(writer, "%s %s %d/%d/%d\n", nome, prestito, day.giorno, day.mese, day.anno)
	}
	writer.Flush()
}

func main() {
	var y int

	MenuInizio()
	fmt.Scanln(&y)

	for y < 1 || y > 3 {
		fmt.Println("Reinserire: ")
		fmt.Scanln(&y)
	}

	for y != 0 {
		switch y {
		case 1:
			LetturaFile()
			fmt.Println("")
		case 2:
			RitiroLibro()
			fmt.Println("")
		case 3:
			RestituzioneLibro()
			fmt.Println("")
		}

		MenuBiblioteca()
		fmt.Scanln(&y)

		for y < 0 || y > 3 {
			fmt.Println("Reinserire: ")
			fmt.Scanln(&y)
		}
	}
}
