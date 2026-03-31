package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CalcQuanteMinMax(frase string) (quante, min, max int){
	if frase == ""{
		return 
	}
	
	var currentLen int
	fraseDivisa := strings.Split(frase, " ")
	lenFraseDivisa := len(fraseDivisa)
	for i:=0; i<lenFraseDivisa; i++{
		parola := fraseDivisa[i]
		currentLen = len(parola)
		// controllo punteggiatura
		if parola[currentLen-1] == SEPARATORE_1 || parola[currentLen-1] == SEPARATORE_2{
			parola = parola[:currentLen-1]
			currentLen--	
		}
		// controlla numero
		_, err := strconv.Atoi(parola) // se err == nil => numero
		if err == nil{
			continue		// salta controlli su lunghezza min e max
		}
		// controlla min max
		if currentLen > max {
			max = currentLen
		}
		if currentLen < min || min == 0{
			min = currentLen
		}
		quante++
	}
	return
}

func TrovaNumDopoKeyword(frase, keyWord string) (num int){
	var err error
	fraseDivisa := strings.Split(frase, " ")
	lenFrase := len(fraseDivisa)
	for i, parola := range fraseDivisa{
		if parola == keyWord{
			if i < lenFrase-2{			// len-2 perchè l'ultimo indice utilizzabile è len-1
				daAnalizzare := fraseDivisa[i+1]
				currentLen := len(daAnalizzare)
				// controllo punteggiatura
				if daAnalizzare[currentLen-1] == SEPARATORE_1 || daAnalizzare[currentLen-1] == SEPARATORE_2{
					daAnalizzare = daAnalizzare[:currentLen-1]
				}
				
				num, err = strconv.Atoi(daAnalizzare)
				if err != nil{
					break 
				}

			}else {
				return 0
			}
		}
	}
	return num
}

func Autogramma(frase string) bool{
	// controlla contiene, massima, minima
	nParole, min, max := CalcQuanteMinMax(frase)
	
	declaredParole := TrovaNumDopoKeyword(frase, CONTIENE)
	declaredMax := TrovaNumDopoKeyword(frase, MASSIMA)
	declaredMin := TrovaNumDopoKeyword(frase, MINIMA)

	return nParole == declaredParole && max == declaredMax && min == declaredMin
}

func StampaAnalisiAutogramma(frase string){
	nParole, min, max := CalcQuanteMinMax(frase)
	declaredParole := TrovaNumDopoKeyword(frase, CONTIENE)
	declaredMax := TrovaNumDopoKeyword(frase, MASSIMA)
	declaredMin := TrovaNumDopoKeyword(frase, MINIMA)
	onesto := Autogramma(frase)

	fmt.Println("=== ", frase)
	fmt.Printf("minimo dichiarato %d contro minimo verificato %d\n", declaredMin, min)
	fmt.Printf("massimo dichiarato %d contro massimo verificato %d\n", declaredMax, max)
	fmt.Printf("numero parole dichiarato %d contro numero parole verificato %d\n", declaredParole, nParole)
	if onesto{
		fmt.Println("onesto")
	}else{
		fmt.Println("bugiardo")
	}
}

const CONTIENE = "contiene:"
const MASSIMA = "massima:"
const MINIMA = "minima:"

const SEPARATORE_1 = ','
const SEPARATORE_2 = ':'

func main(){
	if len(os.Args) < 2{
		fmt.Println("file name?")
		return
	}
	fname := os.Args[1]
	fp, err := os.Open(fname)
	if err != nil{
		fmt.Println("file not found")
		return
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan(){
		frase := scanner.Text()
		if frase != ""{
			StampaAnalisiAutogramma(frase)
		}
	}
}