package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MATRICE struct {
	Nome string
	mat  [][]float64
}

func Creariga(s []string) (riga []float64) {
	if len(s) == 0 {
		return
	}

	num, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		return Creariga(s[1:])
	}

	riga = append(riga, num)
	riga = append(riga, Creariga(s[1:])...)
	return
}

func CreaMat() (matrix []MATRICE) {
	var mat MATRICE
	var counter int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), " ")

		if len(riga) == 1 && counter == 0 {
			mat.Nome = riga[0]
		} else if len(riga) == 1 && counter != 0 {
			matrix = append(matrix, mat)
			mat.Nome, mat.mat = riga[0], [][]float64{}
		} else {
			mat.mat = append(mat.mat, Creariga(riga))
		}
		counter++
	}

	matrix = append(matrix, mat)

	return
}

func (m *MATRICE) Sorting() {
	for _, riga := range m.mat {
		sort.Float64s(riga)
	}
}

func Stampa(m MATRICE) (s string) {
	s += m.Nome + "\n"

	for _, riga := range m.mat {
		for _, numero := range riga {
			s += strconv.FormatFloat(numero, 'f', 2, 64) + "\t"
		}
		s += "\n"
	}

	return
}

func main() {
	var print func(MATRICE) string
	print = Stampa

	matrix := CreaMat()

	for _, m := range matrix {
		m.Sorting()
		fmt.Println(print(m))
	}
}
