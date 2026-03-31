package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {

	if nome != "" && anno > 0 && grado > 0 && quantita > 0 {

		return BottigliaVino{nome, anno, grado, quantita}, true
	} else {

		return BottigliaVino{}, false
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {

	stringa := strings.Split(riga, ",")

	if len(stringa) != 4 {

		return BottigliaVino{}, false

	}

	nome := stringa[0]

	anno, _ := strconv.Atoi(stringa[1])

	grado, _ := (strconv.ParseFloat(stringa[2], 64))

	quantita, _ := strconv.Atoi(stringa[3])

	return CreaBottiglia(nome, anno, float32(grado), quantita)

}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {

	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {

	bottiglia, ok := CreaBottigliaDaRiga((bot))

	if ok {

		AggiungiBottiglia(bottiglia, cantina)
	}

}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {

	var c int

	for i := 0; i < len(cantina); i++ {

		if cantina[i].nome == nome {

			c++
		}
	}

	return c
}

func (b BottigliaVino) String() string {

	stringa := fmt.Sprintf("%s,%d,%.1f,%d", b.nome, b.anno, b.grado, b.quantita)

	pos := strings.Index(stringa, ".0,")

	if pos != -1 {
		stringa = stringa[:pos] + "," + stringa[pos+3:]
	}

	return stringa

}

func main() {

	cantina := []BottigliaVino{}

	file, _ := os.Open(os.Args[1])

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		if scanner.Text() != "" {
			AggiungiBottigliaDaRiga(scanner.Text(), &cantina)
		}

	}

	for _, b := range cantina {

		fmt.Println(b.String())

	}

}

/*

Vini
----

Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

- una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
	nome string
	anno int
	grado float32
	quantita int

- una funzione
	CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool)
  che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.

- una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

  e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
  Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

- una funzione
	CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
  che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
  Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

- una funzione
	AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)
  che aggiunge una bottiglia alla lista

- una funzione
	AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)
  che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo (attenzione che è string, nel formato specificato sopra, una riga)

- una funzione
	ContaPerTipo(nome string, cantina []BottigliaVino) int
  che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina

- un **metodo** (da applicare a BottigliaVino)
    String() string
  che restituisce una riga di descrizione della bottiglia nel seguente formato:  nome,anno,grado,quantità
  (cioè ad es. restituisca "Rasol,2018,14,750" per la prima riga dell'esempio sopra

*/
