package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const DISTANZA int = 32

func Minuscolo(testo []string) (testo1 []string) {
	var riga1 string
	for _, riga := range testo {
		for _, elemento := range riga {
			if unicode.IsUpper(elemento) {
				riga1 += string(int(elemento) + DISTANZA)
			} else {
				riga1 += string(elemento)
			}
		}
		testo1 = append(testo1, riga1)
	}
	return
}

func Maiuscolo(testo []string) (testo1 []string) {
	var riga1 string
	for _, riga := range testo {
		for _, elemento := range riga {
			if unicode.IsLower(elemento) {
				riga1 += string(int(elemento) - DISTANZA)
			} else {
				riga1 += string(elemento)
				continue
			}
		}
		testo1 = append(testo1, riga1)
	}
	return
}

func main() {
	var testo []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo = append(testo, scanner.Text())
	}

	fmt.Println(Maiuscolo(testo))
	fmt.Println(Minuscolo(testo))
}
