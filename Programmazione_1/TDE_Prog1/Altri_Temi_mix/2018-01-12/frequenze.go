package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		s = strings.Fields(line)
	}
	mp := frequenze(s)
	fmt.Println(mp)
}

func frequenze(s []string) map[string]int {
	var mp = make(map[string]int)
	var count int
	for i, x := range s {
		count = 1
		for j, y := range s {
			if x == y && i != j {
				count++
			}
		}
		mp[x] = count
	}
	return mp
}
