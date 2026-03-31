package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	var bottiglia BottigliaVino
	if nome != "" && anno > 0 && grado >= 0 && quantita >= 0 {
		bottiglia.nome, bottiglia.anno, bottiglia.grado, bottiglia.quantita = nome, anno, grado, quantita
		return bottiglia, true
	}
	return bottiglia, false
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	parts := strings.Split(riga, ",")
	anno, _ := strconv.Atoi(parts[1])
	grado64, _ := (strconv.ParseFloat(parts[2], 32))
	grado := float32(grado64)
	quantita, _ := strconv.Atoi(parts[3])
	return CreaBottiglia(parts[0], anno, grado, quantita)
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	bottiglia, ok := CreaBottigliaDaRiga(bot)
	if ok {
		AggiungiBottiglia(bottiglia, cantina)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	var conta int
	for _, name := range cantina {
		if name.nome == nome {
			conta++
		}
	}
	return conta
}

func (bot BottigliaVino) String() string {
	return fmt.Sprintf("%v,%v,%v,%v", bot.nome, bot.anno, bot.grado, bot.quantita)
}

func main() {
	var cantina []BottigliaVino
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			AggiungiBottigliaDaRiga(scanner.Text(), &cantina)
		}
	}

	for _, bot := range cantina {
		fmt.Println(bot.String())
	}
}
