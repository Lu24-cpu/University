package main

import "fmt"

func main() {

	var s string

	for {
		fmt.Println("	Inserisci una parola:")
		fmt.Scan(&s)
		slice := []string{}
		if s == "0" {
			fmt.Println("	Fine")
			break
		}
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				slice = append(slice, string(s[i]))
			}
		}
		if len(slice) == 0 {
			fmt.Println("	Non ci sono doppie")

		} else {
			fmt.Println(slice)
		}

	}

}
