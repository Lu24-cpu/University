package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Studente struct {
	Nome string
	voti [][]float64
}

type Classi struct {
	Alunni []Studente
	Classe string
}

func (student *Studente) String() (s string) {
	s += student.Nome + "\n"
	for _, riga := range student.voti {
		for _, voto := range riga {
			s += strconv.FormatFloat(voto, 'f', 2, 64) + " "
		}
		s += "\n"
	}
	return
}

func (studente *Studente) Medie() {
	for i, materia := range studente.voti {
		if len(materia) == 0 {
			continue
		}
		sum := 0.0
		for _, voto := range materia {
			sum += voto
		}
		media := sum / float64(len(materia))
		studente.voti[i] = append(studente.voti[i], media)
	}
}

func Classe(nomefile string) (Scuola []Classi) {
	var Classe Classi
	var Studenti Studente
	var counter int

	file, err := os.Open(nomefile)
	if err != nil {
		fmt.Println(err)
		return []Classi{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		fmt.Println("Ciao")
		riga, err0 := reader.ReadString('\n')
		riga = strings.TrimSpace(riga) // Rimuove spazi e newline

		if err0 != nil {
			break
		}

		if strings.Contains(riga, "Prima") || strings.Contains(riga, "Seconda") || strings.Contains(riga, "Terza") || strings.Contains(riga, "Quarta") || strings.Contains(riga, "Quinta") {
			if counter > 0 {
				fmt.Println("Ciao0", Scuola, Classe)
				Scuola = append(Scuola, Classe)
				counter, Classe = 0, Classi{}
			}
			fmt.Println("Ciao1")
			Classe.Classe = riga
			counter++
			continue
		}

		fmt.Println("Ciao2")
		elementi := strings.Fields(riga) // Rimuove spazi extra

		_, err1 := strconv.ParseFloat(elementi[0], 64)

		if err1 != nil {
			Studenti.Medie()
			Classe.Alunni = append(Classe.Alunni, Studenti)
			Studenti = Studente{Nome: riga}
		} else {
			var votes []float64
			for _, el := range elementi {
				num, _ := strconv.ParseFloat(el, 64)
				votes = append(votes, num)
			}
			Studenti.voti = append(Studenti.voti, votes)
		}
	}

	if len(Classe.Alunni) > 0 {
		Scuola = append(Scuola, Classe)
	}

	return
}

func main() {
	Collegio := Classe(os.Args[1])

	fmt.Println(Collegio)
	for _, Classe := range Collegio {
		fmt.Println(Classe.Classe)
		for _, Student := range Classe.Alunni {
			fmt.Println(Student.String())
		}
	}
}
