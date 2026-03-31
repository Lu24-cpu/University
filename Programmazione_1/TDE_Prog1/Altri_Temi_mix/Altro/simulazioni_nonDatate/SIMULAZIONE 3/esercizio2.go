package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var n, l, h, count, p, line int

	var s string

	fmt.Scan(&n, &l, &h)

	for {

		for i := 0; i < l; i++ {

			let := strconv.Itoa(count)
			s += let

			count++
			if count == 10 {

				count = 0
			}
		}

		if p%2 == 0 {

			fmt.Println(s)
			s = ""
		} else {

			for i := len(s) - 1; i >= 0; i-- {

				fmt.Print(string(s[i]))
			}
			fmt.Println()
			s = ""
		}
		line++

		if line == (2*n)+1 {

			break
		}

		for i := 0; i < h-2; i++ {

			if p%2 == 0 {
				fmt.Print(strings.Repeat(" ", l-1))
				fmt.Println(count)
				count++
				if count == 10 {

					count = 0
				}
			} else {

				fmt.Println(count)

				count++
				if count == 10 {

					count = 0
				}

			}

		}
		line++

		if line == (2*n)+1 {

			break
		}

		p++

	}

}
