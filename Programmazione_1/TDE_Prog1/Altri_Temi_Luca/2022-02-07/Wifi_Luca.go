// Testo esercizio a fine codice

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type WIFI struct {
	ssid       string
	channel    int
	frequency  int
	signal_dBm int
}

func NewWifi(ssid string, channel, frequency, signal_dBm int) (WIFI, bool) {
	var wifi WIFI
	if ssid != "" && ((channel >= 1 && channel <= 14 && frequency >= 2412 && frequency <= 2484) || (channel >= 7 && channel <= 196 && frequency >= 5035 && frequency <= 5980)) && signal_dBm < -20 {
		wifi = WIFI{ssid: ssid, channel: channel, frequency: frequency, signal_dBm: signal_dBm}
		return wifi, true
	}
	return wifi, false
}

func NewWifiDaStringa(riga string) (WIFI, bool) {
	parts := strings.Split(riga, ",")

	channel, _ := strconv.Atoi(parts[1])
	frequency, _ := strconv.Atoi(parts[2])
	signal_dBm, _ := strconv.Atoi(parts[3])

	return NewWifi(parts[0], channel, frequency, signal_dBm)
}

func (w WIFI) String() string {
	signalWatt := ConvertiDBinWatt(w.signal_dBm)
	return fmt.Sprintf("{%v, %v, %v, %v, %v}", w.ssid, w.channel, w.frequency, w.signal_dBm, signalWatt)
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	return math.Pow(10, float64(signal_dBm)/10) / 1000
}

func PiuPotente(elenco []WIFI, banda string) int {
	maxIndex := -1
	maxSignal := -math.MaxInt32
	banda = strings.ToLower(banda)

	for i, wifi := range elenco {
		isValid := false
		if banda == "2ghz" && wifi.frequency >= 2412 && wifi.frequency <= 2484 {
			isValid = true
		} else if banda == "5ghz" && wifi.frequency >= 5035 && wifi.frequency <= 5980 {
			isValid = true
		} else if banda != "2ghz" && banda != "5ghz" {
			isValid = true
		}

		if isValid && wifi.signal_dBm > maxSignal {
			maxSignal = wifi.signal_dBm
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	var elenco []WIFI
	nomefile, banda := os.Args[1], ""

	if len(os.Args) > 2 {
		banda = os.Args[2]
	}

	file, _ := os.Open(nomefile)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wifi, flag := NewWifiDaStringa(scanner.Text())

		if flag {
			elenco = append(elenco, wifi)
		}
	}

	fmt.Println(elenco[PiuPotente(elenco, banda)])
}

/*

Wifi
----

Si scriva un programma (il file deve chiamarsi 'wifi.go') per gestire l'analisi dei segnali wifi di un ambiente.
Il programma dovrà essere dotato delle seguenti:

- una struttura Wifi con campi (dichiarati in quest'ordine):
	ssid      	string
	channel   	int
	frequency 	int
	signal_dBm  int

- una funzione NewWifi(ssid string, channel int, frequency int, signal_dBm int) Wifi,bool
	che, se i valori passati come parametri rispettano le seguenti condizioni:
	- ssid non vuoto
	- channel
		- tra 1 e 14 (compresi) SE la frequency è tra 2412 e 2484 (compresi)
		- OPPURE tra 7 e 196 SE la frequency è tra 5035 e 5980
	- frequency tra 2412 e 2484 OPPURE tra 5035 e 5980 (compresi estremi)
	- signal_dBm minore di -20 (meno venti)
	crea un'istanza di Wifi opportunamente inizializzata e la restituisce insieme a true, altrimenti restituisce una struttura "zero-value" e false

- una funzione NewWifiDaStringa(line string) Wifi,bool
	che costruisce un'istanza della struttura Wifi a partire da una stringa nel formato CSV (in cui i dati sono separati da virgole, vedere il file 'wifi.csv'), stessi vincoli della funzione definita sopra.
	Il formato è esattamente come nel file, non occorre fare controlli sul formato, ma i dati potrebbero essere non accettabili (ad es. un numero di canale non coerente con la frequenza o l'intestazione del CSV).

- un metodo String da applicare a Wifi
	che restituisca una stringa rappresentativa dei valori della struttura, nella forma:
		{ssid, channel, frequency, signal_dBm, signalW}
	Attenzione che c'è un valore in più rispetto ai dati della struct, va aggiunto un valore calcolato: la potenza del segnale in Watt. Il formato di uscita del valore signalW è quello "naturale" del float64 (formato "%v").

- una funzione ConvertiDBinWatt(signal_dBm int) float64
	che prende come parametro la potenza del segnale in dBm (decibel-milliwatts) e calcola la potenza in Watt, la formula di conversione è:
		Watt = (10^(potenza_in_dBm/10)) / 1000
		Nota: il simbolo ^ indica elevamento a potenza (10^2 è 10 alla seconda)

- una funzione PiuPotente(elenco []Wifi, banda string) int
	che restituisce l'indice della struct che rappresenta il wifi con  segnale (dBm) più potente dell'elenco passato come parametro; in funzione del parametro banda viene restituito il più potente fra i segnali a 2GHz (banda="2GHz"), fra quelli a 5GHz (banda="5GHz") o senza distinzione (banda = qualunque altro valore, compresa la stringa vuota)

- una funzione main()
	che elabora un file (CSV) il cui nome è passato come primo argomento (della linea di comando) e che stampa i dati del segnale più potente sulla base del secondo argomento: se NON presente o se diverso da "5GHz"/"2GHz", trova il più potente senza distinzione di banda; se invece il secondo argomento è presente, stampa i dati secondo la banda richiesta.
	Non occorre fare controlli sulla presenza degli argomenti sulla linea di comando e sul file, potete assumere che il programma venga sempre lanciato correttamente e che il file indicato sia presente e nel formato previsto (vedi sopra).

Esempi di esecuzione
---------------------

$ ./wifi wifi.csv
{at-wind,11,2462,-41,7.94328234724282e-08}

$ ./wifi wifi.csv 5GHz
{at-wind,108,5540,-42,6.30957344480193e-08}

$ ./wifi wifi.csv 2GHz
{at-wind,11,2462,-41,7.94328234724282e-08}

*/
