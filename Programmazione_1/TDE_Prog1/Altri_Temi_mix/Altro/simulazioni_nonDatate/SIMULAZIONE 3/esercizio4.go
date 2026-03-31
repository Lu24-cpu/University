package main

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Punto struct {
	e    string
	x, y int
}
type TriangoloRettangolo struct {
	p1, p2, p3 Punto
}

func Distanza(p1, p2 Punto) float64 {

	return math.Sqrt((((math.Abs(float64(p2.x - p1.x))) * (math.Abs(float64(p2.x - p1.x)))) + ((math.Abs(float64(p2.y - p1.y))) * (math.Abs(float64(p2.y - p1.y))))))

}

func StringPunto(p Punto) string {

	return fmt.Sprintf("%s = (%d, %d)", p.e, p.x, p.y)
}

func StringTriangoloRettangolo(t TriangoloRettangolo) string {

	slice := []float64{Distanza(t.p1, t.p2), Distanza(t.p2, t.p3), Distanza(t.p3, t.p1)}

	sort.Float64s(slice)

	area := (slice[0] * slice[1]) / 2.0

	return fmt.Sprintf("Triangolo rettangolo con vertici %s, %s, %s, ed area %.1f", StringPunto(t.p1), StringPunto(t.p2), StringPunto(t.p3), area)
}
func creaPunto(stringa string) Punto {

	s := strings.Split(stringa, ";")

	x, _ := strconv.Atoi(s[1])
	y, _ := strconv.Atoi(s[2])

	return Punto{s[0], x, y}
}

func lati(t TriangoloRettangolo) ([]float64, float64) {

	l1 := Distanza(t.p1, t.p2)
	l2 := Distanza(t.p2, t.p3)
	l3 := Distanza(t.p3, t.p1)

	slice := []float64{l1, l2, l3}

	sort.Float64s(slice)

	area := (slice[0] * slice[1]) / 2.0

	return slice, area
}
func èRettangolo(t TriangoloRettangolo) bool {

	slice, _ := lati(t)

	l1, l2, l3 := slice[0], slice[1], slice[2]

	if int(((l1 * l1) + (l2 * l2))) == int((l3 * l3)) {

		return true
	}

	return false
}

func main() {

	var s string
	var t, tmax TriangoloRettangolo

	slice := []string{}

	var flag bool

	var max float64

	for {

		_, err := fmt.Scan(&s)

		if err == io.EOF {

			break
		}

		slice = append(slice, s)
	}

	//creo tutti i triangoli possibili

	for i := 0; i < len(slice); i++ {

		for j := i + 2; j < len(slice); j++ {

			//creo i punti per poi creare il triangolo e controllare che sia rettangolo, se lo è controllo il valore dell'area

			p1 := creaPunto(slice[i])
			p2 := creaPunto(slice[i+1])
			p3 := creaPunto(slice[j])

			t = TriangoloRettangolo{p1, p2, p3}

			_, area := lati(t)

			if èRettangolo(t) {

				// se entro nel ciclo ho trovato almeno un triangolo rettangolo

				flag = true

				if area > max {

					max = area

					tmax = t
				}
			}

		}

	}

	// stampo il triangolo con area maggiore solo se è presente un triangolo rettangolo

	if flag == true {

		fmt.Println(StringTriangoloRettangolo(tmax))
	}

}
