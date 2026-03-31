package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type PRODOTTO struct {
	Nome string
	Codice string
}

type MAGAZZINO map[PRODOTTO]int

func NuovoMagazzino() MAGAZZINO {
	Magazzino := make(MAGAZZINO)
	return Magazzino
}

func StringProdotto(prodotto PRODOTTO) string{
	var stringa string
	stringa += "Codie: " + prodotto.Codice + ", Prodotto: " + prodotto.Nome
	return stringa
}

func CreaProdotto(nome string, codice string) PRODOTTO {
	var prodotto PRODOTTO

	prodotto.Nome = nome
	prodotto.Codice = codice
	
	return prodotto
}

func NumeroProdotti(m MAGAZZINO) int {
	var count int
	for _, _ :=range m {
		count++
	}
	return count
}

func ModificaProdotto(m MAGAZZINO, p PRODOTTO, variazione int)(MAGAZZINO, bool){
	var flag1, flag2 bool = false, true
	for i, prodotti := range m.prodotto {
		if strings.Contains(prodotti.Nome, p.Nome){
			m.quantità[i] += variazione
			if m.quantità[i]>0{
				flag1 = true
				break
			} else if m.quantità[i] == 0 {
				var temp MAGAZZINO
				temp.prodotto, m.prodotto = m.prodotto[i+1:], m.prodotto[:i]
				temp.quantità, m.quantità = m.quantità[i+1:], m.quantità[:i]
				m.prodotto = append(m.prodotto, temp.prodotto...)
				m.quantità = append(m.quantità, temp.quantità...)
				flag1 = true
			} else if m.quantità[i] < 0 {
				fmt.Println("L'operazione non è stata svolta a causa dell'insufficenza di prodotti")
				m.quantità[i] -= variazione
				flag1 = false
			}
			flag2 = true
		}else {
			flag2 = false
		}
	}

	if !flag2 && variazione > 0 {
		m.prodotto = append(m.prodotto, p)
		m.quantità = append(m.quantità, variazione)
		flag1 = true
	}

	return m, flag1
}

func main(){
	var lines []string


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	
	Magazzino := NuovoMagazzino()
	for _, prodotto := range lines{
		parts := strings.Split(prodotto, ";")
		variazione, _ := strconv.Atoi(parts[2])
		Magazzino[CreaProdotto(parts[0], parts[1])] += variazione
	}

	fmt.Println("Nel magazzino ci sono:", NumeroProdotti(Magazzino), "prodotti:")
	for key, quantita := range Magazzino {
		stringa := StringProdotto(key)
		fmt.Println(stringa, "Quantità:", quantita)
	}

	Magazzino, flag := ModificaProdotto(Magazzino, Prodotto, variazione)

	if flag {
		fmt.Println("Variazione avvenuta con successo")
	}else {
		var i int
		for i = range Magazzino.prodotto {
			if strings.Contains(Magazzino.prodotto[i].Nome, Prodotto.Nome) {
				break
			}
		}
		fmt.Println("Errore alla riga: ", i+1, StringProdotto(Magazzino.prodotto[i]), "disponibilità: ", Magazzino.quantità[i], "richiesta: ", variazione)
	}
}