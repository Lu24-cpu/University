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

	return math.Sqrt(((p2.y - p1.y) * (p2.y - p1.y)) + ((p2.x - p1.x) * (p2.x - p1.x)))
}

func String(p Punto) string {

	return fmt.Sprintf("%s = (%.1f, %.1f)\n", p.n, p.x, p.y)
}

func Lunghezza(tragitto []Punto) float64 {

	var current, previous Punto
	var lung float64

	current = tragitto[0]

	for i := 1; i < len(tragitto); i++ {

		previous = current
		current = tragitto[i]

		lung += Distanza(previous, current)
	}

	return lung
}

func main() {

	tragitto := NuovoTragitto()

	var dist float64

	fmt.Printf("Lunghezza percorso: %.3f\n", Lunghezza(tragitto))

	for i := 0; i < len(tragitto)-1; i++ {

		dist += Distanza(tragitto[i], tragitto[i+1])

		if dist > Lunghezza(tragitto)/2 {

			fmt.Print(String(tragitto[i+1]))
			break

		}
	}

}
