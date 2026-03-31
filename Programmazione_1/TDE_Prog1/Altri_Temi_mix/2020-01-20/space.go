package main

import "fmt"

type Point3 struct {
	x, y, z float64
}

func proiezione(p *Point3) {

	p.z = 0
}

func punti2D(slice []Point3) (s []Point3) {

	for _, i := range slice {

		if i.z == 0 {

			s = append(s, i)
		}
	}

	return s
}

func main() {
	p1 := Point3{1.3, 2.5, -2}
	p2 := Point3{-3, -2.5, -2}
	proiezione(&p1)
	fmt.Println(p1)
	s := []Point3{p1, p2, {}}
	fmt.Println("punti con z 0:", punti2D(s))

}
