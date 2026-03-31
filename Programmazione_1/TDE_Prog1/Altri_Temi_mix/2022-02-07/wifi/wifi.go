package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

const F_BANDA_1_0 = 2412
const F_BANDA_1_1 = 2484
const F_BANDA_2_0 = 5035
const F_BANDA_2_1 = 5980

const NOME_BANDA_1 = "2GHz"
const NOME_BANDA_2 = "5GHz"

const C_BANDA_1_0 = 1
const C_BANDA_1_1 = 14
const C_BANDA_2_0 = 7
const C_BANDA_2_1 = 196

const SIGNAL_THRESHOLD = -20

type Wifi struct{
	ssid      	string
	channel   	int
	frequency 	int
	signal_dBm  int
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi,bool){
	zero_value_wifi := Wifi{".",0,0,0}
	if ssid == ""{
		return zero_value_wifi, false
	}
	if frequency <= F_BANDA_1_1 && frequency >= F_BANDA_1_0{	// dentro banda frequenze 1
		if channel > C_BANDA_1_1 || channel < C_BANDA_1_0{	// fuori banda channel 1
			return zero_value_wifi, false
		}	
	}else if frequency <= F_BANDA_2_1 && frequency >= F_BANDA_2_0{	// dentro banda frequenze 2
		if channel > C_BANDA_2_1 || channel < C_BANDA_2_0{			// fuori banda channel 2
			return zero_value_wifi, false
		}	
	}
	if signal_dBm >= SIGNAL_THRESHOLD{
		return zero_value_wifi, false
	}
	
	return Wifi{ssid, channel, frequency, signal_dBm}, true
}

func NewWifiDaStringa(line string) (Wifi,bool){
	record := strings.Split(line, ",")
	ssid := record[0]
	channel, _ := strconv.Atoi(record[1])
	frequency, _ := strconv.Atoi(record[2])
	signal_dBm, _ := strconv.Atoi(record[3])

	return NewWifi(ssid, channel, frequency, signal_dBm)
}

func (wifi Wifi) String() string{
	signalW := ConvertiDBinWatt(wifi.signal_dBm)
	return fmt.Sprintf("{%s,%d,%d,%d,%v}",wifi.ssid, wifi.channel, wifi.frequency, wifi.signal_dBm, signalW)
}

func ConvertiDBinWatt(signal_dBm int) float64{
	// Watt = (10^(potenza_in_dBm/10)) / 1000
	watt := (math.Pow(10, float64(signal_dBm)/10)) / 1000
	return watt
}
func returnPotenteEdIndice(current, potente Wifi, potenteIndice, currentIndice int)(Wifi, int){
	if current.signal_dBm > potente.signal_dBm || potente.signal_dBm == 0{
		return current, currentIndice
	}
	return potente, potenteIndice
}
func PiuPotente(elenco []Wifi, banda string) int{
	var potente Wifi
	potente = Wifi{"ssid",0,0,0}
	var indexPotente int
	
	l := len(elenco)
	for i:=0; i<l; i++{
		// controlla in due blocchi se il wifi fa parte di banda 1 o 2, controllando il più potente solo se la banda inserita lo permette
		if elenco[i].frequency <= F_BANDA_1_1 && elenco[i].frequency >= F_BANDA_1_0{	// dentro banda frequenze 1
			if elenco[i].channel <= C_BANDA_1_1 && elenco[i].channel >= C_BANDA_1_0{	// dentro banda channel 1
				if banda == NOME_BANDA_1 || banda != NOME_BANDA_2{
					potente, indexPotente = returnPotenteEdIndice(elenco[i], potente, indexPotente, i)
				}
			}		
		}	

		if elenco[i].frequency <= F_BANDA_2_1 && elenco[i].frequency >= F_BANDA_2_0{	// dentro banda frequenze 2
			if elenco[i].channel <= C_BANDA_2_1 && elenco[i].channel >= C_BANDA_2_0{	// dentro banda channel 2
				if banda == NOME_BANDA_2 || banda != NOME_BANDA_1{
					potente, indexPotente = returnPotenteEdIndice(elenco[i], potente, indexPotente, i)
				}
			}		
		}	
	}
	return indexPotente
}

func readWifis(fname string) (elenco []Wifi){
	fp, _ := os.Open(fname)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	for scanner.Scan(){
		nuovo, ok := NewWifiDaStringa(scanner.Text())
		if ok{
			elenco = append(elenco, nuovo)
		}
	}
	return
}

func main(){
	var elencoWifi []Wifi
	elencoWifi = readWifis(os.Args[1])
	parametroBanda := ""
	if len(os.Args) == 3{
		parametroBanda = os.Args[2]
	}

	piu_potente := PiuPotente(elencoWifi, parametroBanda)
	fmt.Println(elencoWifi[piu_potente])
}