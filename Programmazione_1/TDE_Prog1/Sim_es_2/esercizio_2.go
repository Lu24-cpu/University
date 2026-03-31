package main

import (
	"fmt"
	"os"
	"strconv"
	"sort"
)

func GeneraNumeri(N, K string) (sl []int){
	k, _ := strconv.Atoi(K)
	for i :=range N {
		if i+k+1>len(N) {
			break
		} else {
			numstr := N[:i+1] + N[k+i+1:]
			num, _ := strconv.Atoi(numstr)

			sl = append(sl, num)
		}
		if i==0{
			numstr := N[k:]
			num, _ := strconv.Atoi(numstr)

			sl = append(sl, num)
		}
	}
	return
}

func FiltraNumeri(sl []int)(fl []int){
	for _, numero := range sl {
		flag := 0
		for i:=1; i<numero; i++ {
			if numero % i == 0 {
				flag += i
			}
		}

		if flag > numero {
			fl = append(fl, numero)
		}
	}
	return
}

func main(){
	N, K := os.Args[1], os.Args[2]

	sl := GeneraNumeri(N, K)
	fil := FiltraNumeri(sl)

	sort.Ints(fil)
	for _, numero := range fil{
		fmt.Println(numero)
	}
}