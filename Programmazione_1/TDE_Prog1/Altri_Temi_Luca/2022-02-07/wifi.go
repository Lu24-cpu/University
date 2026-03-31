package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Wifi struct {
	ssid                           string
	channel, frequency, signal_dBm int
}

func NewWifi(ssid string, channel, frequency, signal_dBm int) (Wifi, bool) {
	var wf Wifi
	if len(ssid) != 0 {
		if (channel >= 1 && channel <= 14) && (frequency >= 2412 && frequency <= 2484) {
			if signal_dBm < -20 {
				wf.ssid, wf.channel, wf.frequency, wf.signal_dBm = ssid, channel, frequency, signal_dBm
				return wf, true
			}
		}
		if (channel >= 7 && channel <= 196) && (frequency >= 5035 && frequency <= 5980) {
			if signal_dBm < -20 {
				wf.ssid, wf.channel, wf.frequency, wf.signal_dBm = ssid, channel, frequency, signal_dBm
				return wf, true
			}
		}
	}
	return wf, false
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	parts := strings.Split(line, ",")
	ssid := parts[0]
	channel, _ := strconv.Atoi(parts[1])
	frequency, _ := strconv.Atoi(parts[2])
	signal_dBm, _ := strconv.Atoi(parts[3])
	return (NewWifi(ssid, channel, frequency, signal_dBm))
}

func (wf Wifi) String(signalW float64) string {
	s := fmt.Sprintf("%v, %v, %v, %v, %v", wf.ssid, wf.channel, wf.frequency, wf.signal_dBm, signalW)
	return s
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	signalW := (math.Pow10((signal_dBm / 10))) / 1000
	return signalW
}

func main() {
	var rete []Wifi
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wf, ok := NewWifiDaStringa(scanner.Text())
		if ok {
			rete = append(rete, wf)
		}
	}
	var potenza float64
	var wfPotente Wifi = rete[0]
	potenza = ConvertiDBinWatt(rete[0].signal_dBm)
	if len(os.Args) == 2 {
		for _, wifi := range rete {
			if ConvertiDBinWatt(wifi.signal_dBm) > potenza {
				potenza = ConvertiDBinWatt(wifi.signal_dBm)
				wfPotente = wifi
			}
		}
	} else {
		potenza, _ := strconv.ParseFloat(string(os.Args[2][0]), 64)
		var banda []Wifi
		switch potenza {
		case 2:
			for _, wifi := range rete {
				if wifi.frequency >= 2412 && wifi.frequency <= 2484 {
					banda = append(banda, wifi)
				}
			}
			potenza = ConvertiDBinWatt(banda[0].signal_dBm)
			for _, wifi := range banda {
				if ConvertiDBinWatt(wifi.signal_dBm) > potenza {
					potenza = ConvertiDBinWatt(wifi.signal_dBm)
					wfPotente = wifi
				}
			}
		case 5:
			for _, wifi := range rete {
				if wifi.frequency >= 5035 && wifi.frequency <= 5980 {
					banda = append(banda, wifi)
				}
			}
			potenza = ConvertiDBinWatt(banda[0].signal_dBm)
			for _, wifi := range banda {
				if ConvertiDBinWatt(wifi.signal_dBm) > potenza {
					potenza = ConvertiDBinWatt(wifi.signal_dBm)
					wfPotente = wifi
				}
			}
		}
	}
	fmt.Println(wfPotente.String(potenza))
}
