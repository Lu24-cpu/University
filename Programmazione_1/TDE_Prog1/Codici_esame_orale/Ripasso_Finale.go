/*
Riempimento di matrici ricorsivo e interruzione tramite carattere utilizzando anche i funtion types e specificare il tipo matrice
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MATRICE [][]int

func Riga() (riga []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	el := scanner.Text()

	num, err := strconv.Atoi(el)

	if err != nil {
		return
	}

	riga = append(riga, num)
	riga = append(riga, Riga()...)
	return
}

func Matrice(Riga func() []int) (m MATRICE) {
	var s string

	fmt.Println("Inserire + per una nuova riga - per terminare")
	fmt.Scan(&s)

	if s == "-" {
		return
	}

	m = append(m, Riga())
	m = append(m, Matrice(Riga)...)
	return
}

func (m MATRICE) Stringa() (s string) {
	for _, riga := range m {
		for _, el := range riga {
			s += strconv.Itoa(el) + " "
		}
		s += "\n"
	}
	return
}

func main() {
	var mat MATRICE
	var f func() []int
	f = Riga

	mat = Matrice(f)

	fmt.Println(mat.Stringa())
}
