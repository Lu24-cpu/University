package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Wordle []string = []string{"apice", "sara", "luca", "acqua", "amici", "amicizia", "amore", "anatomia", "andare", "ansia", "arte", "autobus", "bacio", "bagno", "bello", "bere", "biglietto", "bisogno", "blu", "borsa", "caffe", "camera", "cane", "canzone", "capelli", "capire", "casa", "cellulare", "cena", "chiedere", "cibo", "colazione", "computer", "cravatta", "credere", "cucina", "cultura", "domani", "dormire", "fame", "famiglia", "fare", "freddo", "gelosia", "giallo", "gioco", "giornale", "giorno", "internet", "lavare", "lavoro", "leggere", "letto", "lezione", "libro", "macchina", "mamma", "mangiare", "mare", "mattina", "messaggio", "montagna", "morte", "musica", "neve", "noia", "notte", "odio", "oggi", "orario", "papa", "pasta", "peccato", "penna", "pensare", "poesia", "pranzo", "pregare", "prete", "pulire", "regola", "religione", "rosso", "scarpe", "scrivere", "sete", "sigaretta", "soldi", "sonno", "sorella", "studiare", "strumento", "telefono", "televisione", "tempo", "tesoro", "treno", "tutto", "università", "uscire", "verde", "volere", "volontà"}

func EstrazioneParola() (parola string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Intn(len(Wordle))
	parola = Wordle[x]
	return
}

func main() {
	var tentativo string
	var flag1 bool = false
	parola := EstrazioneParola()

	fmt.Println("La lunghezza della parola è:", len(parola), "lettere (inserire una parola tutta minuscola)")

	for i := 0; i < len(parola)+1; i++ {
		fmt.Scan(&tentativo)
		fmt.Println(tentativo, parola)

		if parola == tentativo {
			flag1 = true
			fmt.Println("Coomplimenti hai vinto")
			break
		} else {
			for k := range parola {
				if parola[k] == tentativo[k] {
					fmt.Println("La lettera", string(tentativo[k]), "è nella posizione giusta")
				} else {
					for j := 0; j < len(parola); j++ {
						if parola[j] == tentativo[k] {
							fmt.Println("La lettera", string(tentativo[k]), "è nella posizione sbagliata")
							break
						} else if j == len(parola)-1 {
							fmt.Println("La lettera", string(tentativo[k]), "non è presente nella parola")
						}
					}
				}
			}
		}
		fmt.Println("Nuovo tentativo:")
	}

	if !flag1 {
		fmt.Println("La parola era:", parola)
	}
}
