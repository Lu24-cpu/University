package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	giorno, mese, anno int
}

func Filtro() {
	var stringa string
	var n, k int

	fmt.Println("Inserire una stringa:")
	fmt.Scan(&stringa)

	fmt.Println("Inserire la dimensione nxn:")
	fmt.Scan(&n)

	caratteri := []rune(stringa)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == (n/2) || i == n-1 {
				if k == len(caratteri) {
					k = 0
				}

				fmt.Printf("%c", caratteri[k])
				k++

			} else if j == 0 {
				if k == len(caratteri) {
					k = 0
				}

				if i >= (n/2) && i != n-1 {
					for l := 0; l < n-1; l++ {
						fmt.Print(" ")
					}
				}
				fmt.Printf("%c", caratteri[k])
				k++
			}
		}

		fmt.Println()
	}

	return
}

func Es1() {
	var numeri []int

	fmt.Println("Inserire dei numeri:")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numeri = append(numeri, n)
	}

	fmt.Println(MaxLen(Crescenti(SottoSeq(numeri))))
}

func SottoSeq(num []int) (seq [][]int) {
	for i, el := range num {
		var temp []int
		temp = append(temp, el)
		for j := i + 1; j < len(num); j++ {
			temp = append(temp, num[j])
			if len(temp) >= 2 {
				seq = append(seq, temp)
			}
		}
	}

	return
}

func Crescenti(sottoseq [][]int) (crescenti [][]int) {
	var flag bool = true

	fmt.Println()
	for _, seq := range sottoseq {
		for i, el := range seq {
			if i != 0 && el < seq[i-1] {
				flag = false
			}
		}
		if flag && len(seq) >= 2 {
			crescenti = append(crescenti, seq)
		}
		flag = true
	}

	return
}

func MaxLen(cresce [][]int) (result []int) {
	var max int

	for i, seq := range cresce {
		if i == 0 {
			max = i
		}

		if len(seq) > len(cresce[max]) {
			result = seq
		}
	}

	return
}

func Es2() {
	var data Data

	fmt.Println("Inserire una data (gg, mese, anno):")
	fmt.Scan(&data.giorno, &data.mese, &data.anno)

	anni := data.anno - 1970
	mesi := data.mese - 1
	giorni := data.giorno

	fmt.Println("Sono passati", anni, "anni", mesi, "mesi e", giorni, "giorni")
}

func main() {
	Filtro()
	Es1()
	Es2()
}
