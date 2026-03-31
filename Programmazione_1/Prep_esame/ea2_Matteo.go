package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Persona struct {
	nome, cognome, nickname string
}

func CreaName(s string) Persona {
	var p Persona
	var str string
	for i, v := range s {
		if i == 0 || i == len(s)-1 {
			continue
		}
		str += string(v)
	}
	sl := strings.Split(str, ",")
	p.nome = sl[0]
	p.cognome = sl[1]
	for _, v := range sl {
		for i, k := range v {
			if i == 3 {
				break
			}
			p.nickname += string(k)
		}
	}

	return p
}

func CaratteriDiversi(n1, n2 string) bool {
	for _, v := range n1 {
		for _, k := range n2 {
			if v == k {
				return false
			}
		}
	}
	return true
}

func main() {
	p := make(map[string][]string)
	var testo, n, a string
	var sl, nk []string
	var count int
	var flag bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo = scanner.Text()
		sl = append(sl, testo)
	}
	for _, v := range sl {
		var noco []string
		noco = append(noco, CreaName(v).nome)
		noco = append(noco, CreaName(v).cognome)
		//fmt.Println(noco)
		n = CreaName(v).nickname
		for {
			for _, v := range nk {
				if n != v {
					flag = false
				} else {
					flag = true
					count++
					break
				}
			}
			if !flag {
				break
			} else {
				if count != 1 {
					for j, k := range n {
						if j == len(n)-1 {
							s := strconv.Itoa(count)
							a += s
							break
						}
						a += string(k)
					}
					n = a
					a = ""
					flag = false
					break
				}
				s := strconv.Itoa(count)
				n += s
				flag = false
			}
		}
		nk = append(nk, n)
		count = 0
		fmt.Println(noco, n)
		p[n] = noco
		for k := range p {
			fmt.Println(k, p[k])
		}
	}
	sort.Strings(nk)
	fmt.Println(nk)
}
