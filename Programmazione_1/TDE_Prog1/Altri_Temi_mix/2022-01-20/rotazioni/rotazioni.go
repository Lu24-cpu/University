package main

import (
	"fmt"
)

type Point struct{
	x, y float64
}

type Rectangle struct{
	pLL, pUR Point
}

func NewPoint(x, y float64) Point{
	return Point{x, y}
}

func NewRectangle(p1, p2 Point) Rectangle{
	pLL := Point{min(p1.x, p2.x), min(p1.y, p2.y)}
	pUP := Point{max(p1.x, p2.x), max(p1.y, p2.y)}
	return Rectangle{pLL, pUP}
}

func Rotate(r *Rectangle, dir byte){
	var pLL, pUR Point
	if dir == LEFT{
		pLL = Point{r.pLL.x - (r.pUR.y - r.pLL.y), r.pLL.y}
		pUR = Point{r.pLL.x, r.pLL.y + (r.pUR.x - r.pLL.x)}
	}else if dir == RIGHT{
		pLL = Point{r.pLL.x, r.pLL.y - (r.pUR.x - r.pLL.x)}
		pUR = Point{r.pLL.x + (r.pUR.y - r.pLL.y), r.pLL.y}
	}else{
		return
	}
	r.pLL = pLL
	r.pUR = pUR
}

func (rect Rectangle) String()string{
	// ((P1.x,P1.y) (P3.x,P3.y))
	return fmt.Sprintf("((P1.%f,P1.%f) (P3.%f,P3.%f))", rect.pLL.x, rect.pLL.y, rect.pUR.x, rect.pUR.y)
} 
func max(n1, n2 float64) float64{
	if n1 > n2{
		return n1
	}
	return n2
}
func min(n1, n2 float64) float64{
	if n1 < n2{
		return n1
	}
	return n2
}


const LEFT = 'L'
const RIGHT = 'R'

func main(){
	
}