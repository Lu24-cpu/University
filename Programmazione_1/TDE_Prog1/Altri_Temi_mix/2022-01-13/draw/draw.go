package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func DrawPoint(c byte, k int) string{
	return strings.Repeat(" ", k)+string(c)
}
func DrawSegment(c byte, k, l int) string{
	return strings.Repeat(" ", k)+strings.Repeat(string(c), l)
}

func main(){
	lunghezza, _ := strconv.Atoi(os.Args[1])
	numGradini, _ := strconv.Atoi(os.Args[2])
	carattere := []byte(os.Args[3])[0]		// prendo il primo presupponendo un carattere ASCII

	if lunghezza <= 0 || numGradini <= 0{
		return
	}
	// lunghezza e altezza l, numero n
	for i := 0; i < numGradini; i++{
		fmt.Println(DrawSegment(carattere, i*lunghezza, lunghezza))
		for j:=0; j<lunghezza-1; j++{
			fmt.Println(DrawPoint(carattere, (i+1)*lunghezza-1))
		}
	}


}