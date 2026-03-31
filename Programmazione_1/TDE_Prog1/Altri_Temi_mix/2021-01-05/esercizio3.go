package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	n    string
	x, y float64
}

func NuovoTragitto() (tragitto []Punto) {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		riga := strings.Split(scanner.Text(), ";")

		x, _ := strconv.ParseFloat(riga[1], 64)
		y, _ := strconv.ParseFloat(riga[2], 64)

		tragitto = append(tragitto, Punto{riga[0], x, y})

	}

	return tragitto

}

func Distanza(p1, p2 Punto) float64 {

	return math.Sqrt(((p2.x - p1.x) * (p2.x - p1.x)) + ((p2.y - p1.y) * (p2.y - p1.y)))
}

func String(p Punto) string {

	return fmt.Sprintf("%s = (%.1f, %.1f)\n", p.n, p.x, p.y)

}

func Lunghezza(tragitto []Punto) float64 {

	distanza := 0.0

	var current, previous Punto

	current = tragitto[0]

	for _, i := range tragitto {

		previous = current

		current = i

		distanza += Distanza(previous, current)

	}

	return distanza
}

func main() {

	tragitto := NuovoTragitto()

	fmt.Printf("Lunghezza percorso : %.3f\n", Lunghezza(tragitto))

	meta := Lunghezza(tragitto) / 2.0

	lunghezzaOra := 0.0

	for i := 1; i < len(tragitto)-1; i++ {

		lunghezzaOra += Distanza(tragitto[i], tragitto[i-1])

		if lunghezzaOra > meta {

			fmt.Print("Punto oltre metà: ")

			fmt.Print(String(tragitto[i]))
		}

	}

}
