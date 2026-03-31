package main

import (
	"fmt"
	"os"
	"strconv"
)

type MATRICE [][]int

func CreaMatr(dim int) (mat1, mat2 MATRICE) {
	for i := 0; i < dim; i++ {
		var riga []int
		for j := 0; j < dim; j++ {
			riga = append(riga, (i+1)*(j+1))
		}
		mat1 = append(mat1, riga)
	}

	for i := dim; i >= 0; i-- {
		var riga []int
		for j := dim; j >= 0; j-- {
			riga = append(riga, (i+1)*(j+1))
		}
		mat2 = append(mat2, riga)
	}
	return
}

func SommaMat(i, j, dim int, mat1, mat2, mats MATRICE) MATRICE {
	if i == dim {
		return mats
	}
	if j == dim {
		return SommaMat(i+1, 0, dim, mat1, mat2, mats)
	}

	mats[i][j] = mat1[i][j] + mat2[i][j]

	return SommaMat(i, j+1, dim, mat1, mat2, mats)
}

func main() {
	dim, _ := strconv.Atoi(os.Args[1])
	mat1, mat2 := CreaMatr(dim)

	mats := make(MATRICE, dim)
	for i := range mats {
		mats[i] = make([]int, dim)
	}

	mats = SommaMat(1, 0, dim, mat1, mat2, mats)

	fmt.Println("Matrice somma:")
	for _, righe := range mats {
		for _, elemento := range righe {
			fmt.Print(elemento, " ")
		}
		fmt.Println()
	}
}
