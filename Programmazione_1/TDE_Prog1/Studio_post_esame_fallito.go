package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MATRICE [][]int

func EsFiltro() {
	var elementi []rune

	stringa := os.Args[1]

	elementi = []rune(stringa)

	for i, elemento := range elementi {
		if i < (len(elementi)/2) || i > (len(elementi)/2) {
			for j := 0; j < (len(elementi) / 2); j++ {
				fmt.Print(" ")
			}
			fmt.Printf("%c", elemento)
			for j := 0; j < (len(elementi) / 2); j++ {
				fmt.Print(" ")
			}
			fmt.Println()
		} else if i == (len(elementi) / 2) {
			for j := len(elementi) - 1; j >= 0; j-- {
				fmt.Printf("%c", elementi[j])
			}
			fmt.Println()
		}
	}
}

func EsFiltro2() {
	var elementi []rune

	stringa := os.Args[1]

	elementi = []rune(stringa)

	fmt.Println("")

	for i := range elementi {
		if i < len(stringa)/2 {
			for k := 0; k < (len(stringa)/2)-i; k++ {
				fmt.Print(" ")
			}
			fmt.Printf("%c\n", elementi[len(stringa)/2-i:len(stringa)/2+i+1])
		} else if i > len(stringa)/2 {
			var k int
			for k = 0; k < i-len(stringa)/2; k++ {
				fmt.Print(" ")
			}
			fmt.Printf("%c\n", elementi[k:len(stringa)-k])
		} else {
			fmt.Printf("%c\n", elementi)
		}
	}
}

func Es1Main() {
	var count, r, c int
	var s [][]int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if count == 0 {
			r, _ = strconv.Atoi(parts[0])
			c, _ = strconv.Atoi(parts[1])
		} else {
			var arr []int
			for i := range len(parts) {
				a, _ := strconv.Atoi(parts[i])

				arr = append(arr, a)
			}
			s = append(s, arr)
		}
		count++
	}

	A, errmaines1 := CreaMat(s)

	if errmaines1 == nil {
		Aout, errmaines2 := Rimuovi(A, r, c)

		if errmaines2 == nil {
			fmt.Println("Matrice originale:")
			for _, righe := range A {
				for _, elemento := range righe {
					fmt.Print(elemento, "\t")
				}
				fmt.Println()
			}

			fmt.Println("\nMatrice modificata:")

			for _, righe := range Aout {
				for _, elemento := range righe {
					fmt.Print(elemento, "\t")
				}
				fmt.Println()
			}

		} else {
			fmt.Println(errmaines2)
		}
	} else {
		fmt.Println(errmaines1)
	}
}

func CreaMat(s [][]int) (A MATRICE, errcreamat error) {
	for i, riga := range s {
		if i < len(s)-1 && len(s[i+1]) != len(riga) {
			errcreamat = fmt.Errorf("Le righe sono di lunghezza diversa")
			return
		} else {
			A = append(A, riga)
			errcreamat = nil
		}
	}
	return
}

func Rimuovi(A MATRICE, r, c int) (Aout MATRICE, errimuovi error) {
	for i, riga := range A {
		if len(riga) < c || c < 0 {
			errimuovi = fmt.Errorf("Numero colonne troppo piccolo")
			return
		} else {
			errimuovi = nil
		}

		if i == r {
			continue
		} else {
			var rigaout []int
			for j, colonna := range riga {
				if j == c {
					continue
				} else {
					rigaout = append(rigaout, colonna)
				}
			}
			Aout = append(Aout, rigaout)
		}
	}
	return
}

func main() {
	EsFiltro()
	EsFiltro2()
	Es1Main()
}
