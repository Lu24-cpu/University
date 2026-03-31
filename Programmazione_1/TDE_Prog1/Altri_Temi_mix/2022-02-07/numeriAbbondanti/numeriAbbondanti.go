package main

import (
	"fmt"
)

func Abbondante(n int) bool{
	var somma int
	
	if n <= 0{
		return false;
	}
	for i:=1; i < n/2+1; i++{
		if n % i == 0{	// se è un divisore:
			somma += i
		}
	}
	if n < somma{
		return true
	}
	return false
}

func main(){
	var stampati int
	var daStampare int

	fmt.Scan(&daStampare)
	for i:=1; stampati < daStampare; i++{
		if Abbondante(i){
			fmt.Println(i)
			stampati++
		}
	}
}