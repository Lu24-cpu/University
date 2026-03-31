package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, file := os.Args[1], os.Args[2]
	var testo []string
	fileap, _ := os.Open(file)
	defer fileap.Close()
	scanner := bufio.NewScanner(fileap)
	for scanner.Scan() {
		testo = append(testo, scanner.Text())
	}
	N, _ := strconv.Atoi(n)
	if len(testo) <= N {
		for _, riga := range testo {
			fmt.Println(riga)
		}
	} else {
		for i := (len(testo) - N); i < len(testo); i++ {
			fmt.Println(testo[i])
		}
	}
}
