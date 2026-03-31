package main

import "fmt"

type Data struct {
	gg, mm, aa uint
}

func creaData(g, m, a uint) (Data, bool) {

	if g >= 01 && g <= 30 && a > 0 && (m == 04 || m == 06 || m == 9 || m == 11) {

		return Data{g, m, a}, true
	} else if g >= 01 && g <= 31 && a > 0 && (m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12) {

		return Data{g, m, a}, true
	} else if g >= 01 && g <= 28 && a > 0 && m == 2 {

		return Data{g, m, a}, true
	}

	return Data{0, 0, 0}, false

}

func aggiorna(data *Data) {

	data.gg += 1

	if data.mm == 2 {

		if data.gg == 29 {

			data.gg = 1

			data.mm += 1
		}
	} else if data.mm == 4 || data.mm == 6 || data.mm == 9 || data.mm == 11 {

		if data.gg == 31 {

			data.gg = 1

			data.mm += 11
		}

	} else {

		if data.gg == 32 {

			data.gg = 1
			data.mm += 1
		}
	}

	if data.mm == 13 {

		data.mm = 1
		data.aa += 1
	}
}
func palindroma(d Data) bool {

	var data string

	data += fmt.Sprintf("%.02d%.02d%d", d.gg, d.mm, d.aa)

	j := len(data) - 1

	for i := 0; i < (len(data) / 2); i++ {

		if data[i] != data[j] {

			return false
		}

		j--
	}

	return true

}

func main() {

	var (
		g, m, a uint
		d       Data
		ok      bool
	)
	g, m, a = 1, 2, 2020
	d, ok = creaData(g, m, a)
	if ok {
		fmt.Println("creata:", d)
		aggiorna(&d)
		fmt.Println("data aggiornata:", d)
	}
	if palindroma(d) {
		fmt.Println(d, "data palindroma")
	}

}
