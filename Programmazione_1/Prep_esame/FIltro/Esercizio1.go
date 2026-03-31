package main

import (
	"fmt"
	"os"
	"strconv"
)

func LanciaGenerica(num string) {
	for i := 0; i < len(num)-1; i++ {
		if num[i] < num[i+1] {
			carattere, _ := strconv.Atoi(string(num[i]))
			fmt.Print(carattere)
		}
	}
	fmt.Println()
}

func main() {
	num := os.Args[1]

	LanciaGenerica(num)
}
