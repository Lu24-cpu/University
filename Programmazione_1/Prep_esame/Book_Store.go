package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Cost(books []int) (final int) {
	for _, costo := range books {
		final += costo
	}
	return
}

func main() {
	var libri []string
	var serielibri [][]string
	var costo []float64
	var costo1 []int
	var max int
	sconto := make(map[string]int)
	serie := make(map[string]bool)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserire i libri che si vogliono comprare, un libro costa 8€\n-Se prendi due libri diversi della stessa saga hai il 5% Di sconto\n-Se prendi tre libri diversi della stessa saga hai il 10% Di sconto\n-Se prendi quattro libri diversi della stessa saga hai il 20% Di sconto\n-Se prendi tutti e 5 i libri della stessa saga hai il 25% Di sconto\n")

	for scanner.Scan() {
		libri = append(libri, scanner.Text())
	}

	for i, libro := range libri {
		sconto[libro]++
		if i == 0 {
			max = sconto[libro]
		}
		if max < sconto[libro] {
			max = sconto[libro]
		}
	}

	for i := 0; i < max; i++ {
		var sconti []string

		for key, n := range sconto {
			if n > i {
				sconti = append(sconti, key)
			}
		}

		if len(sconti) > 0 && !serie[strings.Join(sconti, " ")] {
			serielibri = append(serielibri, sconti)
			serie[strings.Join(sconti, " ")] = true
		}
	}

	for _, sconti := range serielibri {
		switch len(sconti) {
		case 1:
			costo = append(costo, 8)
		case 2:
			costo = append(costo, 16*(95/100))
		case 3:
			costo = append(costo, 24*(90/100))
		case 4:
			costo = append(costo, 32*(80/100))
		case 5:
			costo = append(costo, 40*(75/100))
		}
	}

	for _, value := range costo {
		costo1 = append(costo1, int(value*100))
	}
	fmt.Println(Cost(costo1))
}

/*
package bookstore

func Cost(basket []int) int {
	var groups, sets [6]int

	// Group the basket
	// ------------------
	// | Book ID | Count |
	// ------------------
	for i := range basket {
		groups[basket[i]]++
	}

	// Arrange groups into counts by set size
	// --------------------
	// | Set Size | Count |
	// --------------------
	for setSize := 0; sets[0] == 0; sets[setSize], setSize = sets[setSize]+1, 0 {
		for i := range groups {
			if groups[i] > 0 {
				setSize++
				groups[i]--
			}
		}
	}

	// Replace each 3set+5set with two 4sets
	// Not a generic solution; only works for discounts hardcoded below
	for sets[3] > 0 && sets[5] > 0 {
		sets[3], sets[5], sets[4] = sets[3]-1, sets[5]-1, sets[4]+2
	}

	// Return the resultant cost
	return 800*sets[1] + (800*2*0.95)*sets[2] + (800*3*0.9)*sets[3] + (800*4*0.80)*sets[4] + (800*5*0.75)*sets[5]
}
*/
