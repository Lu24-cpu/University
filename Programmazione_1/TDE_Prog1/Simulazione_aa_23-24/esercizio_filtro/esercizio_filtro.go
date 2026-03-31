package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	num := os.Args[1]

	for i:=0; i<len(num)-1; i++{
		a, _ := strconv.Atoi(string(num[i]))
		b, _ := strconv.Atoi(string(num[i+1]))

		if a < b {
			fmt.Print(a)
		}
	}
	fmt.Println()
}