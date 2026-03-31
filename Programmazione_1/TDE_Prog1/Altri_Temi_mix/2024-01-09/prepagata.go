package main

import (
	"errors"
	"fmt"
)

type Prepagata struct {
	numero int
	saldo  float64
}

func (prep Prepagata) String() string {
	return fmt.Sprintf("carta n. %d, saldo %.2f", prep.numero, prep.saldo)
}

func ricarica(carta *Prepagata, importo int) {
	if importo >= 0 {
		carta.saldo = carta.saldo + float64(importo)
	}
}

func preleva(carta *Prepagata, importo int) (int, error) {
	if importo < 0 {
		return 0, errors.New("importo non valido")
	}
	if carta.saldo >= float64(importo) {
		carta.saldo = carta.saldo - float64(importo)
		return importo, nil
	}
	return 0, errors.New("saldo insufficiente")
}

const N_CARTA = 1709
const INIT_SALDO = 100.00
const MOSTRA_SALDO = "a"
const RICARICA_CARTA = "b"
const PRELEVA_CARTA = "c"
const ESCI = "e"

func main() {
	var scelta string
	var importo int
	carta := Prepagata{N_CARTA, INIT_SALDO}

	fmt.Println("a. saldo\nb. ricarica\nc. prelievo\ne. esci")
	fmt.Scan(&scelta)
	for scelta != ESCI {
		switch scelta {
		case MOSTRA_SALDO:
			fmt.Println(carta)

		case RICARICA_CARTA:
			fmt.Scan(&importo)

			ricarica(&carta, importo)
			fmt.Println(carta)

		case PRELEVA_CARTA:
			fmt.Scan(&importo)

			_, errore := preleva(&carta, importo)
			if errore != nil {
				fmt.Println(errore)
			} else {
				fmt.Println(carta)
			}
		default:
			fmt.Println("opzione non valida")
		}
		fmt.Scan(&scelta)
	}
	fmt.Println("arrivederci")
}
