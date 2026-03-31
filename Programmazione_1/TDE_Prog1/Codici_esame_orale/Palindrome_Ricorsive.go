package main

import (
	"fmt"
	"os"
)

func Palindroma(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return Palindroma(s[1 : len(s)-1])
}

func main() {
	if Palindroma(os.Args[1]) {
		fmt.Println("La parola", os.Args[1], "è palindroma")
	} else {
		fmt.Println("La parola", os.Args[1], "non è palindroma")
	}
}
