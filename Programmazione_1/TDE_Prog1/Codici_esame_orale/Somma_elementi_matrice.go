package main

import (
	"fmt"
)

type MATRICE [][]int

func SommaElementi(mat1, mat2 MATRICE, i, dim int) int {
	value := SommaValue(mat1[i], mat2[i], 0, dim)
	if i == dim-1 {
		return value
	}
	return value + SommaElementi(mat1, mat2, i+1, dim)
}

func SommaValue(riga1, riga2 []int, j, dim int) int {
	value := (riga1[j] + riga2[j])
	if j == dim-1 {
		return value
	}
	return value + SommaValue(riga1, riga2, j+1, dim)
}

func main() {
	var mat1, mat2 MATRICE

	mat1 = MATRICE{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	mat2 = MATRICE{[]int{9, 8, 7}, []int{6, 5, 4}, []int{3, 2, 1}}

	fmt.Println(SommaElementi(mat1, mat2, 0, 3))
}
