package main

import (
	"fmt"
	"strings"
	"io"
)

const PARI = true
const STOP = "stop"

func main(){
	var input string
	var parola string
	
	_, err := fmt.Scan(&parola)	
	for err != io.EOF{
		if parola == STOP{
			break
		}
		input += " "+parola
		_, err = fmt.Scan(&parola)	
	}

	listaStringhe := strings.Fields(input)
	for _, parola := range listaStringhe{
		l := len(parola)
		if (l%2 == 0) == PARI{
			for i:=l-1; i>=0; i--{
				fmt.Print(string(parola[i]))
			}
			fmt.Println()
		}
	}
}