package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

type BottigliaVino struct{
	nome string
	anno int
	grado float32
	quantita int
}

func (bot BottigliaVino) String() string{
		return fmt.Sprintf("%s,%d,%s,%d", bot.nome, bot.anno, strconv.FormatFloat(float64(bot.grado), 'f', -1,32), bot.quantita)
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool){
	// (campi string diversi da vuoto, campi int e float maggiori di zero)
	if nome != "" && anno > 0 && grado > 0 && quantita > 0{
		return BottigliaVino{nome, anno, grado, quantita}, true
	}
	return BottigliaVino{"",0,0,0}, false
	
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){
	zero_value_bot := BottigliaVino{"",0,0,0}
	if riga == ""{
		return zero_value_bot, false
	}
	record := strings.Split(riga, SEP)
	if len(record) != 4{
		return zero_value_bot, false
	}
	nome := record[0]
	anno, _ := strconv.Atoi(record[1])
	grado, _ := strconv.ParseFloat(record[2], 32)
	quantita, _ := strconv.Atoi(record[3])
	return CreaBottiglia(nome, anno, float32(grado), quantita)

}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino){
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino){
	bottiglia, ok := CreaBottigliaDaRiga(bot)
	if ok{
		AggiungiBottiglia(bottiglia, cantina)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int{
	var conta int
	for _, bottiglia := range cantina {
		if bottiglia.nome == nome{
			conta++
		}
	}
	return conta
}

const SEP = ","

func main(){
	if len(os.Args)< 2{
		return
	}
	fname := os.Args[1]
	var cantina []BottigliaVino
	fp, err := os.Open(fname)
	if err != nil{
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan(){
		riga := scanner.Text()
		AggiungiBottigliaDaRiga(riga, &cantina)
	}
	for _, bot := range cantina{
		fmt.Println(bot)
	}
}