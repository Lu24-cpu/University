package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	var s string
	fmt.Scan(&s)

	slice := []int{}

	for _, i := range s {

		if unicode.IsDigit(i) {

			num, _ := strconv.Atoi(string(i))

			slice = append(slice, num)
		}

	}

	for i := 0; i < len(slice)-1; i++ {

		if slice[i] < slice[i+1] {

			fmt.Println("Sequenza nascosta non ordinata")
			return
		}

	}

	fmt.Println("Sequenza nascosta ordinata")

}
