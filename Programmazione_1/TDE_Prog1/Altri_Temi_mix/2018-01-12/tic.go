package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orario struct {
	sec int
	min int
	ore int
}

func main() {
	s, _ := strconv.Atoi(os.Args[1])
	m, _ := strconv.Atoi(os.Args[2])
	h, _ := strconv.Atoi(os.Args[3])
	n, _ := strconv.Atoi(os.Args[4])
	ok, orario := newOrario(s, m, h)
	if ok {
		for i := 0; i < n; i++ {
			tic(&orario)
		}
		fmt.Println(orario.sec, ":", orario.min, ":", orario.ore)
	} else {
		fmt.Println("Parametri non validi")
	}
}

func tic(orario *Orario) {
	switch {
	case orario.sec < 59:
		orario.sec++
	case orario.sec == 59 && orario.min != 59:
		orario.sec = 0
		orario.min++
	case orario.min == 59 && orario.min != 59:
		orario.min = 0
		orario.ore++
	case orario.sec == 59 && orario.min == 59:
		orario.min = 0
		orario.sec = 0
		orario.ore++
	}
}

func newOrario(s, m, h int) (bool, Orario) {
	var ok bool
	if s >= 0 && s < 60 && m >= 0 && m < 60 && h >= 0 && h < 24 {
		ok = true
	} else {
		ok = false
	}
	var orario Orario
	orario.sec = s
	orario.min = m
	orario.ore = h
	return ok, orario
}

func conv(s []string) []int {
	var s2 []int
	for _, x := range s {
		y, _ := strconv.Atoi(x)
		s2 = append(s2, y)
	}
	return s2
}
