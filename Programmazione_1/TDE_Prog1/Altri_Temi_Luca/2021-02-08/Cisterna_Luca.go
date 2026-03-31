// testo esercizio a fine codice

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type CISTERNA struct {
	x   float64
	y   float64
	z   float64
	lvl float64
}

func Variazione(cisterna *CISTERNA, lt int) error {
	capacitàMax := int(cisterna.x * cisterna.y * cisterna.z / 1000) // Massima capacità in litri
	contenutoAttuale := Contenuto(*cisterna)

	if lt > 0 { // Aggiungere litri
		if lt <= capacitàMax-contenutoAttuale {
			cisterna.lvl += float64(lt) / (cisterna.x * cisterna.y) * 1000
			return nil
		}
		return errors.New("puoi mettere max " + strconv.Itoa(capacitàMax-contenutoAttuale) + " litri")
	} else { // Togliere litri
		lt = -lt
		if lt <= contenutoAttuale {
			cisterna.lvl -= float64(lt) / (cisterna.x * cisterna.y) * 1000
			return nil
		}
		return errors.New("puoi togliere max " + strconv.Itoa(contenutoAttuale) + " litri")
	}
}

func Contenuto(cisterna CISTERNA) int {
	return int(cisterna.x * cisterna.y * cisterna.lvl / 1000) // Converti cm³ a litri
}

func (c CISTERNA) String() (s string) {
	s = "cisterna: " + strconv.FormatFloat(c.x, 'f', 2, 64) + " cm x " + strconv.FormatFloat(c.y, 'f', 2, 64) + " cm x " + strconv.FormatFloat(c.z, 'f', 2, 64) + " cm"
	return
}

func main() {
	var cisterna CISTERNA
	xs, ys, zs := os.Args[1], os.Args[2], os.Args[3]

	x, _ := strconv.ParseFloat(xs, 64)
	y, _ := strconv.ParseFloat(ys, 64)
	z, _ := strconv.ParseFloat(zs, 64)

	if x > 0 && y > 0 && z > 0 {
		cisterna.x, cisterna.y, cisterna.z, cisterna.lvl = x, y, z, 0
		fmt.Printf("%s\nLivello attuale: %.2f cm, litri: %d\n", cisterna.String(), cisterna.lvl, Contenuto(cisterna))
	} else {
		fmt.Println("Tutti gli argomenti devo essere maggiori di 0")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lts := scanner.Text()
		lt, _ := strconv.Atoi(lts)

		if lt == 9999 {
			break
		}

		err := Variazione(&cisterna, lt)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\nLivello attuale: %.2f cm, litri: %d\n", cisterna.String(), cisterna.lvl, Contenuto(cisterna))
		}
	}
}

/*
Cisterna
========


Scrivere un programma 'cisterna.go' dotato di:

- una struttura 'Cisterna' con 3 campi float64 per le tre dimensioni (in cm) (NOTA SERVIRA' SPECIFICARE NOME E ORDINE DEI CAMPI?) e un ulteriore campo float64 che indica il livello (in cm) a cui arriva l'acqua nella cisterna

- una funzione

    variazione(cisterna *Cisterna, lt int) error

  che, se lt è positivo e c'è abbastanza spazio nella cisterna, versa lt litri (***) nella cisterna e restituisce errore nil, altrimenti non versa niente nella cisterna e restituisce un errore (vedi esempio) segnalando la capienza in litri ancora disponibile; se invece lt è negativo, se ci sono abbastanza litri, prende lt litri dalla cisterna e restituisce errore nil, altrimenti non prende niente dalla cisterna e restituisce un errore (vedi esempio) segnalando la disponibilità (in litri).

- una funzione
	func contenuto(cisterna Cisterna) (litri int)
  che restituisce il numero di litri contenuti nella cisterna

- un metodo String() per Cisterna che restituisce una descrizione della cisterna (vedi esempio sotto)

- una funzione main() che legge da linea di comando tre numeri (float64) che rappresantano larghezza, lunghezza e altezza (in quest'ordine) di una cisterna a base rettangolare.
Se i dati sono tutti > 0, crea una cisterna vuota (livello = 0) con quelle dimensioni e la stampa (vedi es.); altrimenti segnala l'errore "tutti gli argomenti devono essere >0" (se anche uno solo è minore o uguale a 0) o "tutti gli argomenti devono essere numerici" (se anche uno solo non è numerico) e termina.
Poi legge da standard input una sequenza di interi, che rappresentano litri, terminata da 9999: quando il numero letto è positivo, viene messo nella cisterna un numero di litri corrispondenti, viceversa quando il numero è negativo, i litri vengono tolti, sempre che ciò sia possibile, altrimenti viene segnalato un errore (senza però terminare).
Dopo ogni variazione effettiva,  il programma stampa lo stato della cisterna (vedi es.).
Se invece la variazione della cisterna non va a buon fine, il programma non ristampa il suo stato,

*** Nota: un litro corrisponde a un decimetro cubo

Suggerimenti:
1. per creare un errore, utilizzare la funzione New del package 'errors' (vedere la documentazione relativa).

Esempio
=======

(nota: per distinguere input da output, l'input è preceduto da >>>, che però non andrà digitato)

$ ./cisterna 50 100 150
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 0.00 cm, litri 0
>>>1000
puoi mettere max 750 litri
>>>500
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 100.00 cm, litri 500
>>>-600
puoi prendere max 500 litri
>>>-300
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 40.00 cm, litri 200
>>>9999
*/
