// Crea una scacchiera con tutti i pezzi e dare la possibilità di muoverli e stampare di nuovo la scacchiera

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	pezzo string
}

type scacchiera struct {
	scacchi [8][8]Coordinate
	morti   []string
}

func (n *scacchiera) StampaScacchiera(b *scacchiera, lm string) string {
	var s string

	for i, riga := range n.scacchi {
		s += strconv.Itoa(i) + " "
		for j, elemento := range riga {
			if elemento.pezzo != "   " && b.scacchi[i][j].pezzo != "   " && lm == "B" {
				n.morti = append(n.morti, elemento.pezzo)
				elemento.pezzo = "   "
				s += b.scacchi[i][j].pezzo
			} else if elemento.pezzo != "   " && b.scacchi[i][j].pezzo != "   " && lm == "N" {
				b.morti = append(b.morti, b.scacchi[i][j].pezzo)
				b.scacchi[i][j].pezzo = "   "
				s += elemento.pezzo
			} else if elemento.pezzo != "   " && b.scacchi[i][j].pezzo == "   " {
				s += elemento.pezzo
			} else if elemento.pezzo == "   " && b.scacchi[i][j].pezzo != "   " {
				s += b.scacchi[i][j].pezzo
			} else {
				s += "   "
			}
			s += " "
		}
		s += "\n"
	}
	s += "  "
	for i := 0; i < 8; i++ {
		s += " " + strconv.Itoa(i) + "  "
	}

	return s
}

func (s scacchiera) MuoviPezzo(casella, casellanext string) scacchiera {
	partnow := strings.Split(casella, " ")
	partnext := strings.Split(casellanext, " ")

	actualx, _ := strconv.Atoi(partnow[0])
	nextx, _ := strconv.Atoi(partnext[0])
	actualy, _ := strconv.Atoi(partnow[1])
	nexty, _ := strconv.Atoi(partnext[1])

	s.scacchi[nexty][nextx].pezzo, s.scacchi[actualy][actualx].pezzo = s.scacchi[actualy][actualx].pezzo, "   "
	return s
}

func NuovoGiocatore() (b, n scacchiera) {
	var pezzi [8]string = [8]string{"t", "c", "a", "q", "r", "a", "c", "t"}

	for i := 0; i < 8; i++ {
		b.scacchi[len(b.scacchi)-1][i].pezzo, b.scacchi[len(b.scacchi)-2][i].pezzo = pezzi[i]+"_b", "p_b"
		n.scacchi[0][i].pezzo, n.scacchi[1][i].pezzo = pezzi[i]+"_n", "p_n"
	}
	for i := 2; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.scacchi[i-2][j].pezzo = "   "
			n.scacchi[i][j].pezzo = "   "
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	b, n := NuovoGiocatore()

	fmt.Println(n.StampaScacchiera(&b, ""), "\n")
	fmt.Println("\n Chi, pos attuale x y, pos next x y")

	for scanner.Scan() {
		fmt.Println("\n Chi, pos attuale x y, pos next x y")
		caselle := scanner.Text()
		cas := strings.Split(caselle, ", ")
		if cas[0] == "B" {
			b = b.MuoviPezzo(cas[1], cas[2])
		} else if cas[0] == "N" {
			n = n.MuoviPezzo(cas[1], cas[2])
		}
		fmt.Println(n.StampaScacchiera(&b, cas[0]), "\n")
		if cas[0] == "k n" {
			fmt.Println("Il bianco vince")
			break
		} else if cas[0] == "k b" {
			fmt.Println("Il nero vince")
			break
		}
	}
}
