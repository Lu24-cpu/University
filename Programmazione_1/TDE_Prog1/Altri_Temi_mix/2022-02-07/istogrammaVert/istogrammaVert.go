package main

import (
	"fmt"
	"os"
	"strconv"
)

const CHAR = "*"

func max(lista []int) (max int){
	for _, value := range lista{
		if value > max{
			max = value
		}
	}
	return max
}


func main(){
	var valori []int
	// acquisisci dati
	for i:=1; i<len(os.Args); i++{
		intVal, _ := strconv.Atoi(os.Args[i])
		valori = append(valori, intVal)
	}
	valoreMax := max(valori)
	l := len(valori)
	// stampa istogramma
	for riga:=valoreMax; riga > 0; riga--{
		for indiceValore := 0; indiceValore < l; indiceValore++{
			if riga<= valori[indiceValore]{
				fmt.Print(CHAR)
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}