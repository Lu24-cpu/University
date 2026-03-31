// data una parola da riga di comando stampare tutti gli anagrammi possibili

package main

import (
	"fmt"
	"os"
)

func CreAnagrammi(parola0, parola1 []rune) (anagrammi []string) {
	if len(parola1) == 0 { // ritorno dell'anagramma
		return []string{string(parola0)}
	}

	for i := 0; i < len(parola1); i++ { // creazione di tutti gli anagrammi tramite due slice di rune (caratteri), che vengono modificate per creare gli anagrammi
		prefisso := append([]rune{}, parola0...)
		word := append([]rune{}, parola1[:i]...)

		prefisso = append(prefisso, parola1[i])
		word = append(word, parola1[i+1:]...)
		anagrammi = append(anagrammi, CreAnagrammi(prefisso, word)...)
	}

	return
}

/* La mia versione di Anagrammi, con dei problemi di perdita di informazioni e aggiunta di altre non presenti
func CreAnagrammi(parola0, parola1 []rune) (anagrammi []string) {
	if len(parola1) == 0 { // ritorno dell'anagramma
		anagrammi = append(anagrammi, string(parola0))
		return
	}

	for i := 0; i < len(parola1); i++ { // creazione di tutti gli anagrammi tramite due slice di rune (caratteri), che vengono modificate per creare gli anagrammi
		var slice []rune
		if i == 0 { // caso base, se siamo in 0
			parola0 = append(parola0, parola1[0])
			fmt.Printf("%v, %v\n", parola0, parola1[1:])

			if len(parola1) == 0 { // ritorno dell'anagramma
				anagrammi = append(anagrammi, string(parola0))
				return
			}

			anagrammi = append(anagrammi, CreAnagrammi(parola0, parola1[1:])...)
			continue
		}
		// tutti gli altri casi
		parola0 = append(parola0, parola1[i])

		if len(parola1) == 0 { // ritorno dell'anagramma
			anagrammi = append(anagrammi, string(parola0))
			return
		}

		slice = append(parola1[:i], parola1[i+1:]...)
		anagrammi = append(anagrammi, CreAnagrammi(parola0, slice)...)
	}

	return
}
*/

func RimuoviDouble(a []string) (anagrammi []string) { // controllo doppioni
	unici := make(map[string]bool)

	for _, parola := range a { // ciclo che controlla tutti gli anagrammi dello slice, se ne trova due uguali, il secondo viene saltato
		if !unici[parola] {
			unici[parola] = true
			anagrammi = append(anagrammi, parola)
		}
	}

	return
}

func main() {
	parola := os.Args[1]
	lettere := []rune(parola)

	/*for i := 0; i < len(lettere); i++ {
		var single []rune
		if i == 0 {
			single := append(single, lettere[0])
			anagrammi = append(anagrammi, CreAnagrammi(single, lettere[1:])...)
			continue
		}
		single = append(single, lettere[i])
		let := append(lettere[:i], lettere[i+1:]...)
		anagrammi = append(anagrammi, CreAnagrammi(single, let)...)
	}*/

	fmt.Println(RimuoviDouble(CreAnagrammi([]rune{}, lettere)))
}
