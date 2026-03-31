package main 

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func AggiornaDispensa(dispensa map[string]float64, filename string) bool{
	fp, err := os.Open(filename)
	if err != nil{
		return false
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Scan()	// rimuove riga di intestazione
	for scanner.Scan(){
		record := strings.Split(scanner.Text(),",")
		peso, _ := strconv.ParseFloat(record[2], 64)

		if record[0] == APPROVIGIONAMENTO{
			dispensa[record[1]] += peso
		}else if record[0] == USO{
			dispensa[record[1]] -= peso
			if dispensa[record[1]] < 0{
				return false
			}
		}
	}
	return true
}

func Rimanenza(dispensa map[string]float64, alimento string) float64{
	valore, presente := dispensa[alimento]
	if !presente{
		return 0
	}
	return valore
}

const APPROVIGIONAMENTO = "approv"
const USO = "uso"

func main(){
	var dispensa map[string]float64
	dispensa = make(map[string]float64)
	csvName := os.Args[1]

	corretto := AggiornaDispensa(dispensa, csvName)
	l:= len(os.Args)
	if l==2{
		if corretto{
			fmt.Println("file corretto")
		}else{
			fmt.Println("file non corretto")
		}
	}else{
		for i:=2; i<l; i++{
			fmt.Printf("%s %.3f\n",os.Args[i], Rimanenza(dispensa, os.Args[i]))
		}
	}
}