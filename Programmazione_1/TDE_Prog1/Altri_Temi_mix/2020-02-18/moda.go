package main

import "fmt"

func moda(slice []int) (s []int) {

	mappa := make(map[int]int)
	var max int

	for _, i := range slice {

		mappa[i]++

		if mappa[i] > max {

			max = mappa[i]
		}
	}

	for chiave, valore := range mappa {

		if valore == max {

			s = append(s, chiave)
		}
	}

	return s

}

func main() {

	s := []int{5, -2, 3, -2, 3, -2, 3}
	fmt.Println(moda(s))

}
