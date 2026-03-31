package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Persona struct {
	id      int
	nome    string
	cognome string
	età     int
}
type InsiemeTre struct {
	p1, p2, p3 Persona
	etàMedia   float64
}

func StringPersona(p Persona) string {

	return fmt.Sprintf("%d - %s %s: %d", p.id, p.cognome, p.nome, p.età)
}

func StringInsiemeTre(i InsiemeTre) string {

	return fmt.Sprintf("Insieme formato da:\n%s;\n%s;\n%s.\nEtà media: %f\n", StringPersona(i.p1), StringPersona(i.p2), StringPersona(i.p3), i.etàMedia)

}

func rimuoviCifra(nList []string) []string {
	var res []string
	for _, n := range nList {
		for i := range n {
			res = append(res, n[:i]+n[i+1:])
		}
	}
	return res
}

func main() {

	file1, err1 := os.Open(os.Args[1])
	file2, err2 := os.Open(os.Args[2])

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	slice := []Persona{}

	var max float64

	var etàMedia float64

	s := "0123456"

	if err1 != nil || err2 != nil {

		fmt.Println(errors.New("errore nell'apertura del file"))
		return
	}

	for scanner1.Scan() {

		campi := strings.Split(scanner1.Text(), ";")

		id, _ := strconv.Atoi(campi[0])
		età, _ := strconv.Atoi(campi[3])

		slice = append(slice, Persona{id, campi[1], campi[2], età})

	}

	res := []string{s}
	for i := 0; i < 4; i++ {
		res = rimuoviCifra(res)
	}

	for _, i := range res {

		uno, _ := strconv.Atoi(string(i[0]))
		due, _ := strconv.Atoi(string(i[1]))
		tre, _ := strconv.Atoi(string(i[2]))

		p1, p2, p3 := slice[uno], slice[due], slice[tre]

		fmt.Println(p1, p2, p3)

		etàMedia = float64((p1.età + p2.età + p3.età) / 3)

		if etàMedia > max {

			max = etàMedia
		}

	}

	for scanner2.Scan(){

		
	}

	fmt.Println(max)

}
