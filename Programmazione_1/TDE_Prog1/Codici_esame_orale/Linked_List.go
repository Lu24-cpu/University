//Svolgere un esercizio sulle liste lincata

package main

import (
	"fmt"
)

type Node struct {
	stazione string
	Next     *Node
}

func AddFirst(first *Node, station string) *Node {
	s := new(Node)
	s.stazione = station
	s.Next = first
	return s
}

func AddNode(first *Node, station string) *Node {
	New := new(Node)
	prev := new(Node)
	t := first

	New.stazione = station

	for t != nil {
		prev = t
		t = t.Next
	}

	prev.Next = New
	return first
}

func NewLine() *Node {
	var selezione int
	Linea := new(Node)

	for selezione != 3 {

		fmt.Println("Inserire:")
		fmt.Println("1) Inizio")
		fmt.Println("2) Fine")
		fmt.Println("3) Uscire")
		fmt.Scan(&selezione)

		switch selezione {
		case 1: // Vai alla funzione aggiungi all'inizio
			var stazione string

			fmt.Print("Inserire la stazione: ")
			fmt.Scan(&stazione)

			Linea = AddFirst(Linea, stazione)
		case 2: // Vai alla funzione aggiungi alla fine
			var stazione string

			fmt.Print("Inserire la stazione: ")
			fmt.Scan(&stazione)

			Linea = AddNode(Linea, stazione)
		case 3:
			break
		default:
			fmt.Println("Inserimento sbagliato")
		}
	}

	return Linea
}

var counter int
var precedente *Node

func RimuoviStazione(first, stazione *Node, posrim int) *Node {
	t := new(Node)
	t = first

	for i := 0; i <= posrim && t != nil; i++ {
		if i == posrim-1 {
			precedente = t
		}
		if i == posrim {
			precedente.Next = t.Next
		}
		t = t.Next
	}

	return first
}

func main() {
	var cartina []*Node
	var selection, posizione int = 1, 0

	for selection != 0 {
		fmt.Println("Inserire 0 per uscire oppure 1 per aggiungere una nuova linea:")
		fmt.Scan(&selection)

		switch selection {
		case 0:
			break
		case 1:
			cartina = append(cartina, NewLine())
		default:
			fmt.Println("Codice errato")
		}

	}

	for _, linea := range cartina {
		first := linea
		for linea != nil {
			fmt.Printf("%s ", linea.stazione)
			linea = linea.Next
		}
		fmt.Println()
		for selection != 1 {
			fmt.Println("Inserire 0 se si vuole rimuove un elemento o 1 se non si vuole rimuovere niente")
			fmt.Scan(&selection)

			switch selection { // Scelta se rimuovere una stazione o meno
			case 1: // Caso di uscita dal codice
				continue
			case 0: // caso di rimozione della stazione e successiva ristampa della linea modificata
				fmt.Println("Inserire la posizione della stazione da rimuovere:")
				fmt.Scan(&posizione)

				var stazione *Node
				stazione = linea
				linea = RimuoviStazione(first, stazione, posizione)

				fmt.Println("Linea Modificata:")

				for first != nil {
					fmt.Printf("%s ", first.stazione)
					first = first.Next
				}
				fmt.Println()
			default:
				fmt.Println("Codice inserito errato")
			}
		}
	}

}

/*
package main

import (
	"fmt"
)

type Node struct {
	Name string
	Next *Node
}

func AggiungiInizio(first *Node, Nuovo *Node) *Node {
	Nuovo.Next = first
	first = Nuovo
	return first
}

func AggiungiFine(first *Node, Nuovo *Node) *Node {
	t := first
	for t.Next != nil {
		t = t.Next
	}
	t.Next = Nuovo
	return first
}

var counter int
var precedente *Node

func TogliElemento(first, t *Node, k int) *Node {
	if k == 0 {
		first = first.Next
		return first
	}
	if counter == k-1 {
		precedente = t
	} else if counter == k+1 {
		precedente.Next = t
		return first
	}
	counter++
	return TogliElemento(first, t.Next, k)
}

func main() {
	var codice, interrompi, i int
	var cartina []*Node
	for interrompi != 1 {
		fmt.Printf("Vuoi inserire una nuova linea?\n0) Si\n1) No\n")
		fmt.Scan(&interrompi)
		switch interrompi {
		case 0:
		//fmt.Println("qui")
		case 1:
			fmt.Println("salvataggio delle modifiche ...")
			continue
		default:
			fmt.Println("codice non valido")
			continue
		}
		first := new(Node)
		cartina = append(cartina, first)
		codice = 0
		for codice != 3 {
			fmt.Printf("Vuoi inserire una tappa all'inizio o alla fine?\n1) inizio\n2) fine\n3) termina\n")
			fmt.Scan(&codice)
			switch codice {
			case 1: //inizio
				var nome string
				t := new(Node)
				fmt.Println("inserisci la nuova tappa")
				fmt.Scan(&nome)
				t.Name = nome
				first = AggiungiInizio(first, t)
			case 2: //fine
				var nome string
				t := new(Node)
				fmt.Println("inserisci la nuova tappa")
				fmt.Scan(&nome)
				t.Name = nome
				first = AggiungiFine(first, t)
			case 3:
				fmt.Println("salvataggio delle modifiche ...")
				cartina[i] = first
				i++
				continue
			default:
				fmt.Println("codice non valido")
				continue
			}
			cartina[i] = first
		}
	}
	codice = 0
	for codice != 1 {
		fmt.Println("Linee create:")
		for _, c := range cartina {
			for c != nil {
				fmt.Print(c.Name, " ")
				c = c.Next
			}
			fmt.Println()
		}
		fmt.Printf("Vuoi eliminare una tappa in una delle linee create?\n1) No\n2) Si\n")
		fmt.Scan(&codice)
		var linea int
		switch codice {
			case 1:
				fmt.Println("salvataggio delle modifiche ...")
				continue
			case 2:
				fmt.Println("quale linea vuoi modificare")
				fmt.Scan(&linea)
				linea--
			default:
				fmt.Println("codice non valido")
				continue
		}
		copia := cartina[linea]
		for copia != nil {
			fmt.Print(copia.Name, " ")
			copia = copia.Next
		}
		fmt.Println()
		fmt.Println("che posizione vuoi togliere?")
		var k int
		copia = cartina[linea]
		fmt.Scan(&k)
		cartina[linea] = TogliElemento(cartina[linea], copia, k)
	}
	for _, c := range cartina {
		for c != nil {
			fmt.Print(c.Name, " ")
			c = c.Next
		}
		fmt.Println()
	}
}
*/
