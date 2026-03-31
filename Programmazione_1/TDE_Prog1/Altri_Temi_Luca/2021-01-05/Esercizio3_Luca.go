// testo esercizio su pdf

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type PUNTO struct {
	etichetta string
	x, y      float64
}

func NuovoTragitto() (tragitto []PUNTO) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		puntostr := scanner.Text()

		parti := strings.Split(puntostr, ";")
		xfl, _ := strconv.ParseFloat(parti[1], 64)
		yfl, _ := strconv.ParseFloat(parti[2], 64)

		tragitto = append(tragitto, PUNTO{etichetta: parti[0], x: xfl, y: yfl})
	}
	return
}

func Distanza(p1, p2 PUNTO) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func (p PUNTO) String() string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.x, p.y)
}

func Lunghezza(tragitto []PUNTO) (dis float64) {
	for i := 0; i < len(tragitto); i++ {
		dis += Distanza(tragitto[i], tragitto[i+1])
		i++
	}
	return
}

func main() {
	tragitto := NuovoTragitto()

	fmt.Printf("Lunghezza del percorso: %.3f\nPunto oltre la metà: %s\n", Lunghezza(tragitto), tragitto[len(tragitto)/2].String())
}
