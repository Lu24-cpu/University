package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
	"strconv"
)

type PERSONA struct{
	nome, cognome, nickname string
}

func CreaNick(nome, cognome string)(nick string){
	nick = nome[:3]+ cognome[:3]
	return
}

func CaratteriDiversi(s1, s2 string)(flag bool){
	for i:=range s1 {
		for j:=range s2 {
			if s1[i] != s2[j]{
				flag = true
			} else {
				flag = false
			}
		}
	}
	return
} 

func main(){
	var gente []PERSONA
	var persona PERSONA
	var nicknames []string

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		riga := scanner.Text()
		parts := strings.Split(riga[1:len(riga)], ",")
		nick := CreaNick(parts[0], parts[1])

		if len(gente)>0{
			persona.nome, persona.cognome, persona.nickname = parts[0], parts[1], nick

			count := 0
			for i:= range gente {
				if strings.Contains(gente[i].nickname, persona.nickname){
					count++
				}
			}

			if count != 0 {
				persona.nickname += strconv.Itoa(count)
			}
			gente = append(gente, persona)
		} else {
			persona.nome, persona.cognome, persona.nickname = parts[0], parts[1], nick
			gente = append(gente, persona)
		}
	}

	for i:=range gente {
		nicknames = append(nicknames, gente[i].nickname)
	}

	sort.Strings(nicknames)
	fmt.Println(nicknames)

	for k:= range nicknames {
		count := 0
		for j := k+1; j<len(nicknames); j++{
			if !CaratteriDiversi(nicknames[k], nicknames[j]){
				count++
			} else {
				count = 0
				break
			}
		}
		if count == len(nicknames)-k {
			for i:=range gente {
				if nicknames[k] == gente[i].nickname {
					fmt.Println(nicknames[i], "(", gente[i].nome, ",", gente[i].cognome, ")")
					break
				}
			}
		}
	}
}