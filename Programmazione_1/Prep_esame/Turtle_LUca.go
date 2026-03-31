package main

import (
	"image"
	"math"

	"github.com/holizz/terrapin"
)

func Curve(t *terrapin.Terrapin, len float64, level int) {
	if level == 0 {
		t.Forward(len)
	} else {
		Curve(t, len, level-1)
		t.Left(math.Pi / 3.0)
		Curve(t, len, level-1)
		t.Right(2.0 * math.Pi / 3.0)
		Curve(t, len, level-1)
		t.Left(math.Pi / 3.0)
	}
}

func Retti(t *terrapin.Terrapin, len float64, level int) {
	Curve(t, len, level)
	t.Right(2.0 * math.Pi / 3.0)
	Curve(t, len, level)
	t.Right(2.0 * math.Pi / 3.0)
	Curve(t, len, level)
	t.Right(2.0 * math.Pi / 3.0)
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 500, 500))
	t := terrapin.NewTerrapin(i, terrapin.Position{250.0, 450.0})
	Retti(t, 10.0, 3)
}
