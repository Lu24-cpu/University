package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func CalcoloRo(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func CalcoloAlpha(ro, a, b float64) (alpha string) {
	x := a / ro
	switch x {
	case math.Sqrt(3) / 2:
		if b/ro == (1 / 2) {
			alpha = "1/6" + "π"
		} else if b/ro == -(1 / 2) {
			alpha = "-1/6" + "π"
		}
	case 0:
		alpha = "1/2" + "π"
	case math.Sqrt(2) / 2:
		if b/ro == math.Sqrt(2)/2 {
			alpha = "1/4" + "π"
		} else if b/ro == -math.Sqrt(2)/2 {
			alpha = "-1/4" + "π"
		}
	case 0.5:
		if b/ro == math.Sqrt(3)/2 {
			alpha = "1/3" + "π"
		} else if b/ro == -math.Sqrt(3)/2 {
			alpha = "-1/3" + "π"
		}
	case -math.Sqrt(3) / 2:
		if b/ro == (1 / 2) {
			alpha = "5/6" + "π"
		} else if b/ro == -(1 / 2) {
			alpha = "7/6" + "π"
		}
	case -math.Sqrt(2) / 2:
		if b/ro == math.Sqrt(2)/2 {
			alpha = "3/4" + "π"
		} else if b/ro == -math.Sqrt(2)/2 {
			alpha = "5/4" + "π"
		}
	case -0.5:
		if b/ro == math.Sqrt(3)/2 {
			alpha = "2/3" + "π"
		} else if b/ro == -math.Sqrt(3)/2 {
			alpha = "4/3" + "π"
		}
	}

	return
}

func main() {
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)

	ro := CalcoloRo(a, b)
	alpha := CalcoloAlpha(ro, a, b)

	fmt.Printf("%.3f(cos(%s) + isin(%s))\n", ro, alpha, alpha)
}
