package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type UTENZA struct {
	numero_telefono int		
	codice_sim string
}

type RegistroTelefonico map[string][]string

func LeggiUtenze()(utenti []UTENZA){
	var utente UTENZA

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga:=scanner.Text()
		parts := strings.Split(riga, ";")
		utente.numero_telefono, _ = strconv.Atoi(parts[0])
		utente.codice_sim = parts[1]
		utenti = append(utenti, utente)
	}

	return
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utente UTENZA)(registroAggiornato RegistroTelefonico){
	registroAggiornato = registro

	numerostr := strconv.Itoa(utente.numero_telefono)
	registroAggiornato[numerostr] = append(registroAggiornato[numerostr], utente.codice_sim)

	return
}

func Identifica(registro RegistroTelefonico, telefono string)(codiceSIM string){
	x := len(registro[telefono])-1
	codiceSIM = registro[telefono][x]
	return
}

func main() {
	var numero string = "338"
	utenti := LeggiUtenze()
	registro := InizializzaRegistro()

	for i:=range utenti {
		registro = AggiungiUtenza(registro, utenti[i])
	}

	for i :=range registro {
		if strings.Contains(i, numero){
			codiceSIM := Identifica(registro, i)
			fmt.Println("Il numero", i, "è associato al codice Sim più recente è: ", codiceSIM)
		}
	}
}