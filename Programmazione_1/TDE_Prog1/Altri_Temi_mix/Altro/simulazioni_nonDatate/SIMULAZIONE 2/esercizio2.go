package main

import (
	"fmt"
	"os"
	"strconv"
)

func GeneraNumeri(N, k int) []int {

	n := strconv.Itoa(N)
	slice := []int{}

	for i := 0; i < len(n); i++ {

		for j := i; j < len(n); j++ {

			st := n[i : j+1]

			if len(st) == k {

				num, _ := strconv.Atoi(st)

				slice = append(slice, num)
			}

		}
	}

	return slice
}

func FiltraNumeri(sl []int) []int {

	slice := []int{}

	for _, i := range sl {

		if i%2 == 0 {

			slice = append(slice, i)
		}
	}

	return slice
}

func main() {

	N, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])

	slice := (FiltraNumeri(GeneraNumeri(N, n)))

	for _, i := range slice {

		fmt.Println(i)
	}

}
