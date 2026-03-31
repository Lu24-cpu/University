package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Carta struct {
	valore int
	seme   string
}

type Giocatore struct {
	carte []Carta
	prese [][5]Carta
	punti int
}

func CreaCarta(semi string, simboli string) (c Carta) {
	var cardvalue [10]string = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

	c.seme = semi
	for i, el := range cardvalue {
		if simboli == el {
			c.valore = i + 1
		}
	}

	return
}

func CreaMazzo() (Mazzo []Carta) {
	var semi = [4]string{"quadri", "cuori", "picche", "fiori"}
	var simboli = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

	for k := 0; k < 4; k++ {
		for m := 0; m < 10; m++ {
			Mazzo = append(Mazzo, CreaCarta(semi[k], simboli[m]))
		}
	}

	return
}

func Mischia(m *[]Carta) {
	mazzo := *m
	for i := range mazzo {
		j := rand.Intn(i + 1)
		mazzo[i], mazzo[j] = mazzo[j], mazzo[i]
	}

	m = &mazzo
}

func Preleva(m []Carta) (Carta, []Carta) {
	if len(m) == 0 {
		return Carta{}, []Carta{}
	}

	estratta := m[0]

	return estratta, m[1:]
}

func Filegiocatore(player Giocatore, giocatore string) {
	var simboli = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}
	g1, _ := os.Create(giocatore)
	defer g1.Close()

	writer := bufio.NewWriter(g1)

	fmt.Fprintln(writer, "Carte\n")
	for _, carta := range player.carte {
		fmt.Fprintln(writer, simboli[carta.valore-1], carta.seme)
	}

	fmt.Fprintln(writer, "\nPrese:\n")

	for _, presa := range player.prese {
		for _, carta := range presa {
			fmt.Fprint(writer, carta)
			fmt.Fprint(writer, ", ")
		}
		fmt.Fprintln(writer, "\n")
	}
	writer.Flush()
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

func RimozioneCarta(played Carta, carte []Carta) []Carta {
	for k, carta := range carte {
		if carta == played {
			carte = append(carte[:k], carte[k+1:]...)
		}
	}

	return carte
}

func Presa(giocata [4]Carta, seme_briscola string) int {
	var player, numero, prev int
	var seme string
	var briscola bool
	var cardvalue [10]int = [10]int{2, 4, 5, 6, 7, 8, 9, 10, 3, 1}

	//check briscole
	for i, el := range giocata {
		if i == 0 {
			for j, value := range cardvalue {
				if el.valore == value {
					prev = j
				}
			}
		}

		if el.seme == seme_briscola {
			for j, value := range cardvalue {
				if value == el.valore && prev < j {
					player = i
					prev = j
				}
			}
			briscola = true
		}
	}

	//Selezionato il giocatore a cui dare la mano, se presente la briscola semplice, in caso contrario si calcola in base al valore e seme
	if briscola {
		return player
	} else {
		player = 0

		for i, card := range giocata {
			if i == 0 {
				seme, numero = card.seme, card.valore
			}

			if seme != card.seme {
				continue
			} else {
				for j, value := range cardvalue {
					if value == numero {
						numero = j
					}

					if card.valore == value && numero < j {
						numero = j
					}
				}

				player = i
			}
		}

	}

	return player
}

func ContaPunti(giocatori []Giocatore) int {
	var punti0, punti1 int

	punti := SetPunti()

	if len(giocatori) == 2 {
		for _, carta := range giocatori[0].carte {
			punti0 += punti[carta.valore]
		}

		for _, carta := range giocatori[1].carte {
			punti1 += punti[carta.valore]
		}

		if punti0 < punti1 {
			return 1
		} else if punti0 > punti1 {
			return 0
		}

	} else if len(giocatori) == 3 {
		var punti2 int

		for _, carta := range giocatori[0].carte {
			punti0 += punti[carta.valore]
		}

		for _, carta := range giocatori[1].carte {
			punti1 += punti[carta.valore]
		}

		for _, carta := range giocatori[2].carte {
			punti2 += punti[carta.valore]
		}

		if punti0 > punti1 && punti0 > punti2 {
			return 0
		} else if punti1 > punti0 && punti1 > punti2 {
			return 1
		} else if punti2 > punti0 && punti2 > punti1 {
			return 2
		}
	} else if len(giocatori) == 4 {

		for _, carta := range giocatori[0].carte {
			punti0 += punti[carta.valore]
		}

		for _, carta := range giocatori[1].carte {
			punti1 += punti[carta.valore]
		}

		for _, carta := range giocatori[2].carte {
			punti0 += punti[carta.valore]
		}

		for _, carta := range giocatori[3].carte {
			punti1 += punti[carta.valore]
		}

		if punti0 < punti1 {
			return 1
		} else if punti0 > punti1 {
			return 0
		}
	}

	fmt.Println("Partita finita in pareggio")
	return 3
}

func Gioco(nomi []string) {
	var carta, briscola Carta
	var giocata string
	var giocatori [4]Giocatore
	var plaied [4]Carta
	var t int

	mazzo := CreaMazzo()
	Mischia(&mazzo)

	if len(nomi) == 3 {
		duebastoni := Carta{valore: 2, seme: "bastoni"}
		for i, card := range mazzo {
			if card == duebastoni && i != 39 {
				mazzo = append(mazzo[:i-1], mazzo[i:]...)
				break
			} else if card == duebastoni && i == 39 {
				mazzo = mazzo[:i-1]
				break
			}
		}
	}

	for i := 0; i < 3; i++ {
		carta, mazzo = Preleva(mazzo)
		giocatori[0].carte = append(giocatori[0].carte, carta)
		carta, mazzo = Preleva(mazzo)
		giocatori[1].carte = append(giocatori[1].carte, carta)

		if len(nomi) == 3 {
			carta, mazzo = Preleva(mazzo)
			giocatori[2].carte = append(giocatori[2].carte, carta)
		} else if len(nomi) == 4 {
			carta, mazzo = Preleva(mazzo)
			giocatori[2].carte = append(giocatori[2].carte, carta)
			carta, mazzo = Preleva(mazzo)
			giocatori[3].carte = append(giocatori[3].carte, carta)
		}
	}

	briscola, mazzo = Preleva(mazzo)
	mazzo = append(mazzo, briscola)
	fmt.Println("La briscola è", briscola)

	for k := 0; k < 20; k++ {
		for j := 0; j < 4; j++ {
			if len(nomi) == 2 && j == 2 {
				break
			} else if len(nomi) == 3 && j == 3 {
				break
			}

			if t == 4 {
				t = 0
			}

			parts := strings.Split(nomi[t], ".")
			nome := parts[0]
			fmt.Println(nome, "giocare una carta")
			fmt.Scan(&giocata)

			parti := strings.Split(giocata, " ")
			valore, _ := strconv.Atoi(parti[0])
			plaied[j].seme, plaied[j].valore = parti[1], valore

			giocatori[t].carte = RimozioneCarta(plaied[j], giocatori[t].carte)

			Filegiocatore(giocatori[t], nomi[t])

			t++
		}

		for i := 0; i < len(nomi); i++ {
			if len(nomi) == 2 && t == 2 {
				t = 0
			} else if len(nomi) == 3 && t == 3 {
				t = 0
			} else if t == 4 {
				t = 0
			}

			carta, mazzo = Preleva(mazzo)
			giocatori[t].carte = append(giocatori[t].carte, carta)

			t++
		}

		if len(nomi) == 3 {
			k++
		}
		if len(nomi) == 4 {
			k += 2
		}

		//Check presa
		t = Presa(plaied, briscola.seme)

	}

	var giocatorislice []Giocatore
	giocatorislice = append(giocatorislice, giocatori[0])
	giocatorislice = append(giocatorislice, giocatori[1])

	n := ContaPunti(giocatorislice)

	if n != 3 {
		fmt.Println("Ha vinto:", giocatori[n])
	}

	return
}

func main() {
	var nomi []string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Inserire i nomi dei giocatori (Premere ctrl+D per terminare):")

		for scanner.Scan() {
			nomi = append(nomi, scanner.Text())
		}

		if len(nomi) < 2 {
			break
		}

		Gioco(nomi)
	}

}
