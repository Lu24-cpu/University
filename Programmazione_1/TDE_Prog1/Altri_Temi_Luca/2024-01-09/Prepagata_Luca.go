// Testo a fine Esercizio

package main

import (
	"errors"
	"fmt"
)

type PREPAGATA struct {
	numero int
	saldo  float64
}

func (c PREPAGATA) String() string {
	return fmt.Sprintf("carta n. %d, saldo %.2f", c.numero, c.saldo)
}

func ricarica(c *PREPAGATA, value int) {
	if value > 0 {
		c.saldo += float64(value)
	}
}

func prelievo(c *PREPAGATA, value int) (int, error) {
	if float64(value) > c.saldo {
		return 0, errors.New("saldo insufficente")
	}
	if value < 0 {
		return 0, errors.New("inporto non valido")
	}
	c.saldo -= float64(value)
	return value, nil
}

func Menu() {
	fmt.Println("Menu:\na. saldo\nb. ricarica\nc. prelievo\ne. esci")
}

func main() {
	var selection string
	var importo int
	carta := PREPAGATA{1709, 100}

	Menu()
	for selection != "e" {
		fmt.Scan(&selection)

		switch selection {
		case "e":
			fmt.Println("Arrivederci")
			break
		case "a":
			fmt.Println(carta.String())
		case "b":
			fmt.Scan(&importo)
			if importo < 0 {
				fmt.Println("opzione non valida")
			}
		case "c":
			fmt.Scan(&importo)
			_, err := prelievo(&carta, importo)
			if err == nil {
				fmt.Println(carta)
			} else {
				fmt.Println(err)
			}
		default:
			fmt.Println("importo non valido")
		}
	}
}

/*

Carta prepagata
===============

Scrivere un programma 'prepagata.go' dotato di:

- una struttura 'Prepagata' che ha due campi, 'numero' di tipo int e 'saldo' di tipo float64, dichiarati in quest'ordine, che rappresentano il numero della carta prepagata e il saldo disponibile sulla carta stessa

- una funzione

    ricarica(carta *Prepagata, importo int)

  che carica importo sulla carta, aggiungendolo al saldo già disponibile. Se importo è negativo, non fa nulla.


- una funzione

    preleva(carta *Prepagata, importo int) (int, error)

  che preleva importo dalla carta, e restituisce la cifra prelevata e un errore;
  in particolare, se importo è maggiore o uguale a 0 e il saldo è maggiore o uguale all'importo,
  addebita importo sulla carta e restituisce importo e 'nil';
  altrimenti restituisce 0 e un errore "saldo insufficiente" in caso di mancanza di fondi o "importo non valido" in caso di un parametro negativo.


- un *metodo* String per Prepagata che restituisce una descrizione della carta nel seguente formato:

	carta n. <numero>, saldo <saldo>

  con due cifre decimali per il saldo. Ad esempio:

	carta n. 2341, saldo 200.15


- una funzione main() che crea una carta prepagata con numero 1709 e saldo 100, stampa un menu:
a. saldo
b. ricarica
c. prelievo
e. esci

e chiede ripetutamente all'utente un'opzione fino a che l'opzione non è 'e'.

Il main:
- se l'opzione è 'e', stampa "arrivederci" e termina;
- se l'opzione è 'a', stampa la situazione della carta (vedi il formato nelle specifiche di String());
- se l'opzione è 'b', legge da stdin l'importo digitato dall'utente e lo carica sulla carta, poi stampa la situazione della carta;
- se l'opzione è 'c', legge da stdin l'importo digitato dall'utente e lo addebita sulla carta, poi stampa la situazione della carta o un messaggio di errore;
- per ogni altra opzione, stampa "opzione non valida".

Suggerimenti:
1. per creare un errore, utilizzare la funzione New
   del package 'errors' (vedere la documentazione relativa).
2. per concatenare strighe e numeri, utilizzare le funzioni
   Sprint o Sprintf del package fmt (vedere la documentazione relativa).
3. per rappresentare un float con 2 cifre decimali, usare %.2f

Esempio
=======

data la seguente sequenza di opzioni in input:

a
c
200
b
500
c
200
c
-78
b
-86
f
e

$ ./prepagata

produce in output:

a. saldo
b. ricarica
c. prelievo
e. esci
carta n. 1709, saldo 100.00
saldo insufficiente
carta n. 1709, saldo 600.00
carta n. 1709, saldo 400.00
importo non valido
carta n. 1709, saldo 400.00
opzione non valida
arrivederci

*/
