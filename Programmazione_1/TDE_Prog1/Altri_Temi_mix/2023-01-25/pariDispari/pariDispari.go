package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	// true per pari, false per dispari
	var pari bool

	if len(os.Args) < 2 {
		fmt.Println("nessun valore in ingresso")
		return
	}
	for i, value := range os.Args[1:]{
		int_value, err := strconv.Atoi(value)
		if err != nil{
			fmt.Printf("elemento in posizione %d non valido\n", i+1)
			return
		}
		
		if i == 0{
			pari = !(int_value % 2 == 0)
		}else if pari == (int_value % 2 == 0){
			fmt.Printf("elemento in posizione %d non valido\n", i+1)
			return
		}
		pari = !pari
	}
	fmt.Println("sequenza valida")
}