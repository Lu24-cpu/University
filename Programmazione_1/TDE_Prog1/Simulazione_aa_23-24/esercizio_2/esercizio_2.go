package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"unicode"
)

type MAPPA map[string]int

func SottoStringhe(str string) MAPPA {
	mappa := make(MAPPA)
	var flag bool

	for i := 0; i<len(str)-1; i++ {
		if str[i]<str[i+1]{
			flag = true
			mappa[str[i:i+2]]++
			for j:=2; j<len(str)-i; j++{
				if flag && str[i+j-1]<str[j+i] {
					mappa[str[i:i+j+1]]++
				}else {
					flag = false
				}
			}
		}
	}
	return mappa
}

func main(){
	var sottostringhe []string
	var stringa string
	var flag bool = true

	fmt.Println("Inserire una stringa senza spazi")
	fmt.Scan(&stringa)

	for i:=0; i<len(stringa)-1; i++ {
		if unicode.IsUpper(rune(stringa[i])){
			flag = false
		}
	}
	
	if flag== true {
		mappa := SottoStringhe(stringa)

		for key, _ := range mappa {
			sottostringhe = append(sottostringhe, key)
		}	

		sort.Strings(sottostringhe)

		for i := range sottostringhe {
			fmt.Println("Sottostringa:", sottostringhe[i], "\tRipetizioni:", mappa[sottostringhe[i]])
		}
	}else {
		fmt.Println("è stata inserita una lettera maiuscola")
	}
}