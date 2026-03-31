package main

import (
	"fmt"
	"strconv"
)

type Carta struct {
	valore int
	seme   string
}

type Giocatore struct {
	carte []Carta
	prese [][5]Carta
	flag  bool
	punti int
}

func SetPunti() (punti map[int]int) {
	punti = make(map[int]int)

	for i := 1; i <= 10; i++ {
		punti[i] = 0
	}

	punti[1] = 11
	punti[3] = 10
	punti[8] = 2
	punti[9] = 3
	punti[10] = 4

	return
}

func Contapunti(player [5]Giocatore, punto int) (result string) {
	var compagnipunti int
	var punteggio [5]int

	punti := SetPunti()

	for i, g := range player {
		for _, prese := range g.prese {
			for _, carta := range prese {
				for key, vcarta := range punti {
					if carta.valore == key {
						g.punti += vcarta
					}
				}
			}
		}
		punteggio[i] = g.punti
	}

	//Somma punti dei due compagni
	for i, giocatore := range player {
		if giocatore.flag {
			compagnipunti += punteggio[i]
		}
	}

	//Stampa dei vincitori
	if compagnipunti > punto && punto != 0 {
		result = "Hanno vinto i compagni, con: "
		result += strconv.Itoa(compagnipunti)
		return
	}

	if compagnipunti == 60 && punto == 0 {
		result = "La partita è finita in pareggio, complimenti"
		result += strconv.Itoa(compagnipunti)
		return
	}

	if compagnipunti < 60 {
		punteggio := 120 - compagnipunti
		result = "Hanno vinto i buoni con: "
		result += strconv.Itoa(punteggio)
		return
	}

	return
}

func main() {
	var player [5]Giocatore = [5]Giocatore{
		{
			carte: []Carta{},
			prese: [][5]Carta{
				{{2, "Cuori"}, {4, "Picche"}, {6, "Quadri"}, {8, "Fiori"}, {10, "Cuori"}},
				{{3, "Fiori"}, {5, "Quadri"}, {7, "Picche"}, {9, "Cuori"}, {1, "Fiori"}},
			},
			flag:  true,
			punti: 0,
		},
		{
			carte: []Carta{},
			prese: [][5]Carta{
				{{1, "Cuori"}, {3, "Picche"}, {5, "Quadri"}, {7, "Fiori"}, {9, "Cuori"}},
				{{2, "Fiori"}, {4, "Quadri"}, {6, "Picche"}, {8, "Cuori"}, {10, "Fiori"}},
			},
			flag:  false,
			punti: 0,
		},
		{
			carte: []Carta{},
			prese: [][5]Carta{
				{{4, "Fiori"}, {6, "Quadri"}, {8, "Picche"}, {10, "Cuori"}, {2, "Fiori"}},
			},
			flag:  true,
			punti: 0,
		},
		{
			carte: []Carta{},
			prese: [][5]Carta{
				{{3, "Cuori"}, {6, "Picche"}, {9, "Quadri"}, {1, "Fiori"}, {4, "Cuori"}},
				{{5, "Fiori"}, {7, "Quadri"}, {2, "Picche"}, {8, "Cuori"}, {10, "Fiori"}},
			},
			flag:  false,
			punti: 0,
		},
		{
			carte: []Carta{},
			prese: [][5]Carta{
				{{1, "Fiori"}, {3, "Quadri"}, {5, "Picche"}, {7, "Cuori"}, {9, "Fiori"}},
			},
			flag:  false,
			punti: 0,
		},
	}

	fmt.Println(Contapunti(player, 0))
	fmt.Println(Contapunti(player, 75))
	fmt.Println(Contapunti(player, 80))
}
