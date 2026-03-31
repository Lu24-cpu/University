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

const (
	Asso    = 1
	Due     = 2
	Tre     = 3
	Quattro = 4
	Cinque  = 5
	Sei     = 6
	Sette   = 7
	Fante   = 8
	Donna   = 9
	Re      = 10
)

type Seme string

const (
	Quadri     Seme = "quadri"
	Cuori      Seme = "cuori"
	Picche     Seme = "picche"
	Fiori      Seme = "fiori"
	NessunSeme Seme = ""
)

var tuttiISemi = []Seme{Quadri, Cuori, Picche, Fiori}

type chiamata struct {
	Carta Carta
	Seme  Seme
	Passa bool
}

// Questa struct formata solo da una variabile int indica a quali posizioni è l'IA rispetto all'ordine iniziale del giro
type GiocatoreIA struct {
	ID int
}

// Questa struttura definisce la carta, ha una variabile di tipo intero per il valore della carta e il seme della stessa
type Carta struct {
	valore int
	seme   Seme
}

// Struttura che definisce il giocatore, cioè c'è salvata la mano del giocatore, le carte che ha preso in due slice del tipo carta
// poi c'è una flag che indica se è uno dei due compagni (T) o meno (F) e i punti complessivi fatti dal singolo giocatore
type Giocatore struct {
	carte []Carta
	prese [][5]Carta
	flag  bool
	punti int
}

func CreaCarta(semi Seme, simboli string) (c Carta) {
	//Questa funzione prende in input un seme e genera di conseguenza una carta di quel seme convertendo il valore della carta in un intero
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
	//Questa funzione viene chiamata dal main per creare tutto il mazzo, quundi chiama la funzione crea carta che ritorna un elemento di tipo Carta
	var semi = [4]string{"quadri", "cuori", "picche", "fiori"}
	var simboli = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

	for k := 0; k < 4; k++ {
		for m := 0; m < 10; m++ {
			Mazzo = append(Mazzo, CreaCarta(Seme(semi[k]), simboli[m]))
		}
	}

	return
}

func Mischia(m *[]Carta) {
	//Questa funzione usa la pseudo randomicità per mischiare il mazzo creato in precedenza
	mazzo := *m
	for i := range mazzo {
		j := rand.Intn(i + 1)
		mazzo[i], mazzo[j] = mazzo[j], mazzo[i]
	}

	m = &mazzo
}

func Preleva(m []Carta) (Carta, []Carta) {
	//Questa funzione preleva una carta dal mazzo come se stesse distrubuendo le carte, di conseguenza toglie quella carta dal mazzo
	if len(m) == 0 {
		return Carta{}, []Carta{}
	}

	estratta := m[0]

	return estratta, m[1:]
}

// Questa funzione scrive su un file le informazioni di ogni giocatore, le carte in mano, le carte prese durante il gioco e se sei il compagno
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

// Sezione gioco multiplayer senza IA
func Chiamata(player [5]Giocatore) (chiamata Carta, punto int, pl [5]Giocatore) {
	var flag bool = true       //questa flag serve per continuare a ciclare l'inserimento dei punti quando si rimane in 2 al 2
	var carta, briscola string //la variabile carta è la carta che viene chiamata durante "l'asta", mentre briscola è la variabile che indica la carta che viene scelta a fine asta
	var i, count, g1, g2 int   //la variabile i è il counter per i giocatori, la variabile count conta quanti giocatori hanno passato, le variabili g1 e g2 indicano chi sono i due che sono arrivati al 2
	var cardvalue [10]string = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

	passi := make(map[int]bool) //questa mappa serve a tener conto di chi passa e insieme alla variabile I a saltare i giocatori che hanno passato e far chiamare direttamente chi non ha ancora passato
	//La variabile i serve a chiamare il giocatore che deve fare la chiamata e la mappa si ricorda chi abbia passato o meno con una flag
	fmt.Println("Istruzioni\nInserire uno alla volta, dal primo al quinto giocatore la carta chiamata (valore numerico)")
	fmt.Println("Se si vuole passare scrivere passo")
	fmt.Println("Inserire 'sfidaadue' se si vuole decidere i punti a cui arrivare")

	for k := 0; k < 5; k++ {
		passi[k] = false //tutta la mappa viene settata su false e poi ogni volta che un giocatore passa la variabile del giocatore passa su true
	}

	for { //questo for serve a ciclare in continuo finchè non rimangono due giocatori al 2 o uno solo prima del 2
		if i == 5 {
			i = 0 //Resetto a 0 la i se arriva a 5 (visto che i giocatori sono 5 e partono da 0)
		}
		if passi[i] { //Qua controllo se il giocatore ha già passato, in caso affermativo passa al giocatore dopo
			i++
			continue
		}

		nome := strings.Split(os.Args[i+1], ".") //Qui estraggo dall'imput di terminale il nome dell'utente che deve chiamare
		fmt.Println("Giocatore", nome[0])
		fmt.Scan(&carta)

		if !passi[i] && carta == "passo" { //Qua controllo che il giocatore che dovrebbe chiamare non abbia già passato
			passi[i] = true
			count++
		}

		if (carta == "sfidaadue" && count == 3) || (count == 4 && carta == "passo") { //Con questo if controllo se la fase della chiamata è terminata oppure si passa alla sfida a due
			break
		}

		for _, card := range cardvalue {
			if carta == card {
				briscola = carta //Qua controllo e cambio la briscola corrente, questo ciclo converte la stringa della carta inserita nel suo formato di struct Carta
			}
		}
		i++ //Iterazione dei giocatori
	}

	i = 0 //Resetto il contatore a 0 per far si che posso utilizzarlo anche nella sfida a due
	if carta == "sfidaadue" {
		for key, chiama := range passi { //Con questo for mi prendo i due giocatori che hanno chiamato al 2
			if !chiama && i == 0 {
				g1 = key
			} else if !chiama && i != 0 {
				g2 = key
			}
			i++
		}

		i = 0

		for flag { //Con questo ciclo gestisco la sfida dei due giocatori alla chiamata a punti
			fmt.Println("Inserire i punti a cui arrivare (per lasciare inserire una parola qualsiasi):")
			fmt.Scan(&carta) //faccio inserire i punti, ma sono in formato stringa

			num, err := strconv.Atoi(carta) //converto la stringa inserita in un int, per cui se da errore non è un numero, per cui il giocatore ha passato
			if err == nil {
				punto = num //Salvataggio del punteggio di arrivo
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

	for key, chiama := range passi { //Con questo ciclo faccio fare la chiamata del seme al chiamante
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

	for i, el := range cardvalue { //In questo for cerco il valore della carta nell'array cardvalue e lo salvo nella struct Carta nominata briscola
		if el == briscola {
			chiamata.valore = i + 1
			break
		}
	}

	pl = player

	return
}

func Presa(giocata [5]Carta, giocatori []string, seme_briscola Seme) string {
	var player, numero, prev int                                   //la variabile player ricorda l'indice del giocatore che ha preso, la variabile numero indica il valore della carta più alta giocata, la variabile prev indica il valore della penultima carta controllata
	var seme Seme                                                  //questa variabile indica il seme della carta che comanda
	var briscola bool                                              //questa variabile bool indica se sono presenti briscole o meno
	var cardvalue [10]int = [10]int{2, 4, 5, 6, 7, 8, 9, 10, 3, 1} //Semplice array con il valore delle carte

	//check briscole
	for i, el := range giocata {
		if i == 0 { //con questo if controllo la prima carta giocata e salvo il valore della carta giocata
			for j, value := range cardvalue {
				if el.valore == value {
					prev = j
				}
			}
		}

		if el.seme == Seme(seme_briscola) { //In questo caso controllo tutte le carte se c'è una briscola e assegno la posizione della briscola al giocatore che l'ha messa
			for j, value := range cardvalue {
				if value == el.valore && prev < j { //controllo se ci sono altre briscole e se sono più o meno alte
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
	} else { //Se non ci sono briscole allora controllo chi ha preso effettivamente la mano
		player = 0

		for i, card := range giocata { //Giro tutte le carte giocate e ad ogni carta controllo se avendo lo stesso seme hanno valore più alto o meno
			if i == 0 { //alla prima iterazione prendo la carta e me la salvo
				seme, numero = card.seme, card.valore
			}

			if seme != card.seme { //controllo che la carta corrente sia dello stesso seme o meno
				continue
			} else { //In caso la carta corrente sia dello stesso seme della prima inserita allora controllo se è più alta o più bassa
				for j, value := range cardvalue {
					if value == numero {
						numero = j
					}

					if card.valore == value && numero < j {
						numero = j
						player = i //essendo la carta dello stesso seme e salvo il giocatore che sta prendendo
					}
				}

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
			if first > 4 { //questo if resetta il primo giocatore se arriva al sesto
				first = 0
			}

			parti := strings.Split(scanner.Text(), " ") //Qui recupero il valore della carta giocata e il suo seme
			valore, _ := strconv.Atoi(parti[0])
			giocata[t].seme, giocata[t].valore = Seme(parti[1]), valore

			played := giocata[t]
			player[first].carte = played.RimozioneCarta(player[first].carte)

			t++
			if t >= 5 { //Controllo che non stia giocando il sesto giocatore
				break
			}

			first++

			if first > 4 {
				first = 0
			}

			if t <= 4 { //Qui recupero il nome del giocatore che deve buttare la carta
				nome := strings.Split(os.Args[first+1], ".")
				fmt.Println("Giocatore", nome[0], "buttare la carta (valore seme) (ps il valore scritto a numero)")
			}
		}

		t = 0

		pl := append(os.Args[giocatore+1:], os.Args[1:giocatore+1]...) // pl indica i giocatori in modo tale che ci sia prima il giocatore che ha preso l'ultima mano e poi gli altri a turno
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

		nome1 := strings.Split(os.Args[giocatore+1], ".") //una volta salvata sul file del giocatore la mano presa allora viene stampato chi ha preso e il nome del prossimo giocatore
		fmt.Println("Mano presa dal giocatore", nome1[0], "Si va alla prossima")
	}

	return player
}

func (played Carta) RimozioneCarta(carte []Carta) []Carta {
	//Questa funzione prende in input una mano di un giocatore e rimuove la carta giocata da un giocatore
	for k, carta := range carte {
		if carta == played {
			carte = append(carte[:k], carte[k+1:]...)
		}
	}

	return carte
}

func SetPunti() (punti map[int]int) {
	//In questa funzione semplicemente creo una mappa con i punti delle carte per il momento della conta dei punti
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
	//FUnzione necessaria a fine della partita per contare i punti dei giocatori e delle due squadre
	var compagnipunti int
	var punteggio [5]int

	punti := SetPunti()

	for i, g := range player { // Primo for per girare i giocatori
		for _, prese := range g.prese { // Secondo ciclo per leggere tutte le mani prese dal giocatore
			for _, carta := range prese { // Terzo per controllare le carte di ogni mano
				for key, vcarta := range punti { // Quarto per controllare l'effettivo valore delle carte
					if carta.valore == key { // Mi salvo il punteggio
						g.punti += vcarta
					}
				}
			}
		}
		punteggio[i] = g.punti // Salvo il punteggio fatto dal giocatore
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

// Sezione per il gioco per cui sono necessari meno di 5 giocatori fisici ed entra in gioco l'IA
func valutaMano(mano []Carta) int {
	punteggio := 0
	for _, c := range mano {
		punteggio += c.valore // Asso=11, 3=10, etc.
	}
	return punteggio
}

// ritorna vero se è presente una carta nella mano, se no ritorna falso
func haCarta(mano []Carta, valore int, seed Seme) bool {

	for _, carta := range mano {
		if carta.seme == seed && carta.valore == valore {
			return true
		}
	}

	return false
}

// filtra le carte per un certo seme
func filtraCartePerSeme(mano []Carta, seed Seme) []Carta {
	var seme []Carta

	for _, carta := range mano {
		if carta.seme == seed {
			seme = append(seme, carta)
		}
	}

	return seme
}

// Migliorato il funzionamento della scelta della chiamata IA
func scegliChiamata(mano []Carta) (cartaChiamata Carta, semeBriscola Seme) {
	// Per ogni seme controlla la carta che si può chiamare
	for _, seme := range tuttiISemi {
		carteNelSeme := filtraCartePerSeme(mano, seme) //filtraggio della mano per i 4 semi

		if len(carteNelSeme) >= 2 {
			// Caso 1: ho Asso e 3 → provo a chiamare il Re
			if haCarta(mano, Asso, seme) && haCarta(mano, Tre, seme) && !haCarta(mano, Re, seme) {
				return Carta{Re, seme}, seme
			}
			// Caso 2: ho Re e 3 ma mi manca l'Asso → chiamo l'Asso
			if haCarta(mano, Re, seme) && haCarta(mano, Tre, seme) && !haCarta(mano, Asso, seme) {
				return Carta{Asso, seme}, seme
			}
			// Caso 3: ho Asso, Re e 3 → provo a chiamare una carta più bassa
			if haCarta(mano, Asso, seme) && haCarta(mano, Re, seme) && haCarta(mano, Tre, seme) {
				// Provo a chiamare in ordine di forza: Cavallo, Fante, 7, 6, 5, 4, 2
				for _, valore := range []Valore{Cavallo, Fante, Sette, Sei, Cinque, Quattro, Due} {
					if !haCarta(mano, valore, seme) && valore < cartaChiamata.valore {
						return Carta{valore, seme}, seme
					}
				}
			}
		}
	}

	// Nessuna chiamata utile → passo
	return Carta{11, ""}, ""
}

func (g *GiocatoreIA) FaiChiamata(mano []Carta, chiamataCorrente chiamata) chiamata {
	forza := valutaMano(mano)

	if forza < 50 { //Da modificare il valore di forza in base a quanto si vuole fare difficile l'IA
		return chiamata{Passa: true}
	}

	chiamataProposta, seme := scegliChiamata(mano)

	// Confronta con la chiamataCorrente
	if èChiamataPiùDebole(chiamataProposta, chiamataCorrente.Carta) {
		return chiamata{
			Carta: chiamataProposta,
			Seme:  seme,
			Passa: false,
		}
	}

	return chiamata{Passa: true}
}

func èChiamataPiùDebole(chiamataProposta Carta, chiamataCorrente Carta) bool {
	if chiamataProposta.valore > chiamataCorrente.valore {
		return true
	}

	return false
}

func FiltraPerSeme(mano []Carta, seme Seme) (carte []Carta) {
	for _, el := range mano {
		if el.seme == seme {
			carte = append(carte, el)
		}
	}

	return
}

func ChiamataIA(player [5]Giocatore, IA [4]GiocatoreIA) (chiamata Carta, punto int, pl [5]Giocatore) {
	//In questa funzione va implementata la chiamata con giocatori e IA
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

		//controllo se il giocatore è un IA e chiamata IA
		for j := range IA {
			//Se trova l'IA corrispondente breaka e poi bisogna mandare una flag per i giocatori umani?
			if IA[j].ID == i {
				temp, _ := scegliChiamata(player[i].carte)

				if temp.valore > 10 {
					carta = "passo"
				}

				flag = true
				break
			} else {
				flag = false
			}
		}

		if !flag {
			nome := strings.Split(os.Args[i+1], ".")
			fmt.Println("Giocatore", nome[0])
			fmt.Scan(&carta)
		}

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

	//Ancora da implementare la sfida a due tra IA e IA e tra IA e umano
	i, flag = 0, true
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

	//Scelta del seme di briscola, ancora da implementare il fatto che il chiamante sia l'IA
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

func GiocoIA(pl [5]Giocatore, briscola Carta, IA [4]GiocatoreIA) (player [5]Giocatore) {
	// Sarà divertente programmare questa parte...
	player = pl

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
			giocata[t].seme, giocata[t].valore = Seme(parti[1]), valore

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

func main() {
	var player [5]Giocatore
	var briscola Carta
	var count, punto int = 1, 0
	giocatori := os.Args[1:]

	//Creazione del mazzo da gioco e rimescolamento delle carte
	mazzo1 := CreaMazzo()
	Mischia(&mazzo1)

	if len(giocatori) == 5 {

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
		chiamata, punti, p := Chiamata(player)
		punto = punti
		player = p
		//Salvataggio dei due compagni, tramite una variabile booleana settata a True se sono i compagni
		for i, g := range player {
			for _, carta = range g.carte {
				if carta == chiamata {
					g.flag = true
					player[i].flag = true
					g.punti = punti
					Filegiocatore(g, giocatori[i])
				}
			}
		}

		//Inizio fase di gioco
		player = Gioco(chiamata, player)
	} else {
		//To do chiamata func IA (Occhio alle condizioni dell'if)
		var IA [4]GiocatoreIA

		for i := range IA {
			if (5 - len(giocatori)) < i {
				break
			}

			IA[i].ID = i
		}

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

		briscola, punto, player = ChiamataIA(player, IA)

		player = GiocoIA(player, briscola, IA)
	}

	//Conta dei punti
	Contapunti(player, punto)
	// Ciao sto utilizzando il comando cat per modificare il file
}
