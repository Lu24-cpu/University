package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MATRICE [][]int

func Sommamat(mat1, mat2 MATRICE) MATRICE {
	var mats MATRICE

	if len(mat1) == 0 {
		return mat2
	} else if len(mat2) == 0 {
		return mat1
	}

	mats = append(mats, SommaElementi(mat1[0], mat2[0]))
	mats = append(mats, Sommamat(mat1[1:], mat2[1:])...)
	return mats
}

func SommaElementi(riga1, riga2 []int) []int {
	var riga []int

	if len(riga1) == 0 || len(riga2) == 0 {
		return riga
	}

	riga = append(riga, riga1[0]+riga2[0])
	riga = append(riga, SommaElementi(riga1[1:], riga2[1:])...)
	return riga
}

func AssegnaMat1() (mat MATRICE) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Inserire una riga alla volta la prima matrice, ogni elemento separato da uno spazio:")
	for scanner.Scan() {
		var elementi []int
		riga := scanner.Text()
		parts := strings.Split(riga, " ")

		for _, elemento := range parts {
			n, _ := strconv.Atoi(elemento)
			elementi = append(elementi, n)
		}
		mat = append(mat, elementi)
	}

	return
}

func AssegnaMat2() (mat MATRICE) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Inserire una riga alla volta la seconda matrice, ogni elemento separato da uno spazio:")
	for scanner.Scan() {
		var elementi []int
		riga := scanner.Text()
		parts := strings.Split(riga, " ")

		for _, elemento := range parts {
			n, _ := strconv.Atoi(elemento)
			elementi = append(elementi, n)
		}
		mat = append(mat, elementi)
	}

	return
}

func main() {
	var mat1, mat2 MATRICE

	mat1 = AssegnaMat1()
	mat2 = AssegnaMat2()

	mats := Sommamat(mat1, mat2)

	for _, righe := range mats {
		fmt.Println(righe)
	}
}
