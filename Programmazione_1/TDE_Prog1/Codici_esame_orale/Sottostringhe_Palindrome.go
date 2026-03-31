/*
package main

import (
	"fmt"
	"os"
)

func IsPalindrome(s string) (flag bool) {
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-1-i] {
			flag = true
		} else {
			flag = false
			return
		}
	}
	return
}

func Anagrammi(input []rune, parola string, anagrammi []string) []string {
	if len(input) == 0 {
		anagrammi = append(anagrammi, parola)
		return anagrammi
	}

	for i, el := range input {
		anagrammi = Anagrammi(input[i+1:], parola+string(el), anagrammi)
	}
	return anagrammi
}

func OrdinaSottoseq(sottoseq []string, input string) (sequenze []string) {
	for _, parola := range sottoseq {
		var i int
		var flag bool = true
		for _, char := range input {
			if rune(parola[i]) == char && i == 0 {
				i++
				continue
			}

			if rune(parola[i]) == char {
				i++
			} else if i != 0 {
				flag = false
			}
		}
		if flag {
			sequenze = append(sequenze, parola)
		}
	}
	return
}

func Sottosequenze(inizial []rune, sottoseq []string) []string {
	var risposte []string
	for _, anagr := range sottoseq {
		var s string
		var flag bool = false
		var count int
		if IsPalindrome(anagr) {
			for _, char := range inizial {
				if char == rune(anagr[0]) {
					flag = true
				} else if char == rune(anagr[len(anagr)-1]) {
					flag = false
					continue
				}
				if !flag {
					s += string(char) + " "
				} else if count < 1 {
					count++
					s += anagr
				}
			}
			risposte = append(risposte, s)
		}
	}
	return risposte
}

func main() {
	sottoseq := Anagrammi([]rune(os.Args[1]), "", []string{})
	fmt.Println(sottoseq)
	sottoseq = OrdinaSottoseq(sottoseq, os.Args[1])
	fmt.Println(sottoseq)

	palindrome := Sottosequenze([]rune(os.Args[1]), sottoseq)

	for _, parola := range palindrome {
		fmt.Println(parola)
	}
}
*/

// partizioni palindrome di una parola
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var count int
var st []string

func Anagrammi(risposta []string, parola, s string) []string {
	t := []rune(s)
	if len(t) == 0 {
		risposta = append(risposta, parola)
		return risposta
	}
	for i := 0; i < len(t); i++ {
		risposta = Anagrammi(risposta, parola+string(t[i]), string(s[i+1:]))
	}
	return risposta
}

func Palindrome(sottostringhe []string) (pali []string) {
	if len(sottostringhe) == 0 {
		return
	}
	parts := strings.Split(sottostringhe[0], " ")
	var flag bool = true
	for _, prova := range parts {
		if !IsPalindrome(prova) {
			flag = false
		}
	}
	if flag {
		pali = append(pali, sottostringhe[0])
	}

	pali = append(pali, Palindrome(sottostringhe[1:])...)
	return
}

func IsPalindrome(parola string) bool {
	if len(parola) <= 1 {
		return true
	}

	if parola[0] != parola[len(parola)-1] {
		return false
	}

	return IsPalindrome(parola[1 : len(parola)-1])
}

func SottoStringhe(parole []string, parola, inizio string) []string {
	var str string
	var flag bool = false
	var i int
	for k, c := range inizio {
		if c == rune(parola[i]) {
			if !flag && k != 0 {
				str += " "
			}
			str += string(c)
			flag = true
			i++
		} else if flag {
			str += " " + string(c)
			flag = false
		} else if !flag {
			str += string(c)
		}
	}

	st = append(st, str)
	count++
	if count == len(parole) {
		return st
	}
	return SottoStringhe(parole, parole[count], inizio)
}

func main() {
	s := os.Args[1]
	var n string
	for i, _ := range s {
		n += strconv.Itoa(i)
		//Vengono convertiti gli indici della stringa in stringa e messi uno dietro l'altro per evitare confusioni con parole che hanno delle doppie una dietro l'altr
	}
	ris := Anagrammi(nil, "", n)
	sn := SottoStringhe(ris, ris[0], n)
	var sl []string
	var parola string
	for _, val := range sn {
		for _, carattere := range val {
			if unicode.IsSpace(carattere) {
				parola += string(carattere)
			} else {
				n, _ := strconv.Atoi(string(carattere))
				parola += string(s[n])
			}
		}
		sl = append(sl, parola)
		parola = ""
	}
	pl := Palindrome(sl)
	for _, x := range pl {
		fmt.Println(x)
	}
}
