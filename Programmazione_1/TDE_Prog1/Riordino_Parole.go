package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var parole []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parole = append(parole, scanner.Text())
	}

	for i := range parole {
		j := rand.Intn(i + 1)
		parole[i], parole[j] = parole[j], parole[i]
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	for _, parola := range parole {
		fmt.Println(parola)
	}
}
