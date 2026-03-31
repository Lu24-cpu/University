package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ProdottoVettoriale(a, b []int)(mat [][]int){
	for i := range a{
		var riga []int

		for j := range b{
			prod := a[i]*b[j]
			riga = append(riga, prod)
		}
		mat = append(mat, riga)
	}

	return
}

func main(){
	var a, b []int
	var num []string

	fmt.Println("Inserire due array di numeri, intervallati da un invio:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		num = append(num, scanner.Text())
	}

	for i, stringa := range num{
		parts:=strings.Split(stringa, " ")
		if i==0 {
			for i:=range parts{
				numero, _ := strconv.Atoi(parts[i])
				a = append(a, numero)
			}
		}else{
			for i:=range parts{
				numero, _ := strconv.Atoi(parts[i])
				b = append(b, numero)
			}
		}
	}

	mat := ProdottoVettoriale(a, b)

	for _, prod := range mat {
		fmt.Println(prod)
	}
}