package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rightNums []int
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		if s.Text() == "0" {
			break
		}
		num, _ := strconv.Atoi(s.Text())
		if num > 17 && num < 31 {
			rightNums = append(rightNums, num)
		}
	}
	Percentuali(rightNums)
}

func Percentuali(s []int) {
	fmt.Println(s)
	var count1, count2, count3, count4 int
	for _, x := range s {
		switch {
		case x > 17 && x < 22:
			count1++
		case x > 21 && x < 25:
			count2++
		case x > 24 && x < 28:
			count3++
		case x > 27 && x < 31:
			count4++
		}
	}
	tot := len(s)
	fmt.Println("A : ", count4*100/tot, " %	")
	fmt.Println("B : ", count3*100/tot, " %")
	fmt.Println("C : ", count2*100/tot, " %")
	fmt.Println("D : ", count1*100/tot, " %")
}
