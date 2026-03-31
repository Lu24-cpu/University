package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*

Regole Briscola chiamata:
- Prima parte la chiamata della briscola
    1. Si chiama il valore della carta, partendo dall'asso fino ad arrivare al due
    2. Se in due giocatori arrivano a chiamare al due, cioè la carta più debole di tutte, si passa a chiamare il punteggio
        2.1 Se si arriva a chiamare il punteggio si chiama partendo da almeno 61 fino ad un limite in cui il giocatore se la sente
    3. Quando rimane un singolo giocatore che chiama e gli altri passano allora quell'ultimo dice il seme di briscolo

- Seconda parte gioco effettivo
    1. Ad ogni turno ogni giocatore può giocare una sola carta
    2. Per prendere ci sono due modi
        2.1 giocare una carta dello stesso seme della prima e che superi il valore della stessa
        2.2 giocare una carta di briscola (se già presente una vale anche in questo caso la 2.1)
    3. Il primo giocatore a giocare la carta è l'ultimo ad aver preso durante la partita (se è la prima mano solitamente il giocatore dopo il mazziere)

- Terza parte conteggio punti, in questa sezione vengono semplicemtente contati i punti (asso = 11 pt, tre = 10 pt, re = 4 pt, donna = 3 pt, jack = 2 pt.

*/

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
	if player.flag {
		fmt.Fprintln(writer, "Compagno")
	}
	writer.Flush()
}

func Chiamata(player [5]Giocatore) (chiamata Carta, punto int, pl [5]Giocatore) {
	var flag bool = true
	var carta, briscola string
	var i, count, g1, g2 int
	var cardvalue [10]string = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

	passi := make(map[int]bool)

	fmt.Println("Istruzioni\nInserire uno alla volta, dal primo al quinto giocatore la carta chiamata (valore numerico)")
	fmt.Println("Se si vuole passare scrivere passo")
	fmt.Println("Inserire 'sfidaadue' se si vuole decidere i punti a cui arrivare")

	for k := 0; k < 5; k++ {
		passi[k] = false
	}

	for {
		if i == 5 {
			i = 0
		}
		if passi[i] {
			i++
			if passi[i] {
				i++
				continue
			}
		}

		nome := strings.Split(os.Args[i+1], ".")
		fmt.Println("Giocatore", nome[0])
		fmt.Scan(&carta)

		if !passi[i] && carta == "passo" {
			passi[i] = true
			count++
		}

		if (carta == "sfidaadue" && count == 3) || (count == 4 && carta == "passo") {
			break
		}

		for _, card := range cardvalue {
			if carta == card {
				briscola = carta
			}
		}
		i++
	}

	i = 0
	if carta == "sfidaadue" {
		for key, chiama := range passi {
			if !chiama && i == 0 {
				g1 = key
			} else if !chiama && i != 0 {
				g2 = key
			}
			i++
		}

		i = 0

		for flag {
			fmt.Println("Inserire i punti a cui arrivare (per lasciare inserire una parola qualsiasi):")
			fmt.Scan(&carta)

			num, err := strconv.Atoi(carta)
			if err == nil {
				punto = num
			} else {
				if i%2 == 1 {
					passi[g2] = true
				} else {
					passi[g1] = true
				}
				flag = false
			}
			i++
		}
	}

	for key, chiama := range passi {
		if !chiama {
			player[key].flag = true
			player[key].punti = punto
			Filegiocatore(player[key], os.Args[key+1])

			nome := strings.Split(os.Args[key+1], ".")
			fmt.Println("Player", nome[0], "inserire il seme di briscola:")
			fmt.Scan(&chiamata.seme)
			break
		}
	}

	for i, el := range cardvalue {
		if el == briscola {
			chiamata.valore = i + 1
			break
		}
	}

	pl = player

	return
}

func Presa(giocata [5]Carta, giocatori []string, seme_briscola string) string {
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
		return giocatori[player]
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

	return giocatori[player]
}

func Gioco(carta_briscola Carta, player [5]Giocatore) [5]Giocatore {
	var giocatore, t, first int

	scanner := bufio.NewScanner(os.Stdin)

	for k := 0; k < 8; k++ {
		var giocata [5]Carta

		first = giocatore

		//Momento della giocata delle carte
		nome := strings.Split(os.Args[first+1], ".")
		fmt.Println("Giocatore", nome[0], "buttare la carta (valore seme) (ps il valore scritto a numero)")

		for scanner.Scan() {
			if first > 4 {
				first = 0
			}

			parti := strings.Split(scanner.Text(), " ")
			valore, _ := strconv.Atoi(parti[0])
			giocata[t].seme, giocata[t].valore = parti[1], valore

			played := giocata[t]
			player[first].carte = played.RimozioneCarta(player[first].carte)

			t++
			if t >= 5 {
				break
			}

			first++

			if first > 4 {
				first = 0
			}

			if t <= 4 {
				nome := strings.Split(os.Args[first+1], ".")
				fmt.Println("Giocatore", nome[0], "buttare la carta (valore seme) (ps il valore scritto a numero)")
			}
		}

		t = 0

		pl := append(os.Args[giocatore+1:], os.Args[1:giocatore+1]...)
		nomeg := Presa(giocata, pl, carta_briscola.seme)

		// Salvataggio sul file di ogni giocatore
		for i, nome := range os.Args[1:] {
			if nome == nomeg {
				//Assegnazione della mano al giocatore con la carta non di briscola più alta
				player[i].prese = append(player[i].prese, giocata)
				giocatore = i
			}
			Filegiocatore(player[i], nome)
		}

		nome1 := strings.Split(os.Args[giocatore+1], ".")
		fmt.Println("Mano presa dal giocatore", nome1[0], "Si va alla prossima")
	}

	return player
}

func (played Carta) RimozioneCarta(carte []Carta) []Carta {
	for k, carta := range carte {
		if carta == played {
			carte = append(carte[:k], carte[k+1:]...)
		}
	}

	return carte
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

func Contapunti(player [5]Giocatore, punto int) {
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
		fmt.Println("Hanno vinto i compagni, con:", compagnipunti)
	}

	if compagnipunti == 60 && punto == 0 {
		fmt.Println("La partita è finita in pareggio, complimenti")
	}

	if compagnipunti < 60 {
		var punteggio int
		punteggio = 120 - compagnipunti
		fmt.Println("Hanno vinto i buoni con:", punteggio)
		fmt.Println(compagnipunti)
	}
}

func main() {
	var player [5]Giocatore
	var count int = 1
	giocatori := os.Args[1:]

	//Creazione del mazzo da gioco e rimescolamento delle carte
	mazzo1 := CreaMazzo()
	Mischia(&mazzo1)

	//Distribuzione delle carte ai 5 giocatori, una a persona
	carta, mazzo := Preleva(mazzo1)
	player[0].carte = append(player[0].carte, carta)

	for i := 0; i < 40; i++ {
		if count == 5 {
			count = 0
		}
		carta, mazzo = Preleva(mazzo)
		if carta.seme != "" {
			player[count].carte = append(player[count].carte, carta)
		}
		count++
	}

	//Salvataggio delle carte dei giocatori su un file txt il cui nome è passato da linea di comando
	for i, nome := range giocatori {
		Filegiocatore(player[i], nome)
	}

	//Funzione che serve per effettuare le chiamate con ritorno del seme di briscola e della carta chiamata
	chiamata, punto, p := Chiamata(player)

	player = p
	//Salvataggio dei due compagni, tramite una variabile booleana settata a True se sono i compagni
	for i, g := range player {
		for _, carta = range g.carte {
			if carta == chiamata {
				g.flag = true
				player[i].flag = true
				Filegiocatore(g, giocatori[i])
			}
		}
	}

	//Inizio fase di gioco
	player = Gioco(chiamata, player)

	//Conta dei punti
	Contapunti(player, punto)
}
