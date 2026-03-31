package main

import (
	"fmt"
	"os"
	"sort"
	"unicode"
)

func SottoStringhe(stringa string) (st []string, flag bool) {
	flag = true
	for i := 0; i < len(stringa)-1; i++ {
		if stringa[i] == stringa[i+1] {
			for j := i + 1; j < len(stringa); j++ {
				if unicode.IsLower(rune(stringa[j])) {
					flag = false
					break
				} else if stringa[j-1] == stringa[j] {
					st = append(st, stringa[i:j+1])
				} else {
					break
				}
			}
		}
	}
	return
}

func main() {
	stringa := os.Args[1]

	st, flag := SottoStringhe(stringa)

	sort.Strings(st)
	if flag {
		fmt.Println(st)
	}
}
