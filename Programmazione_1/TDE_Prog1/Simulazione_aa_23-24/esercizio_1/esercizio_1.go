package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"math"
)

func main(){
	f := os.Args[1]
	soglia, _ := strconv.ParseFloat(os.Args[2], 64)
	epsilon, _ := strconv.ParseFloat(os.Args[3], 64)

	var a, b, c, num int
	var count int
	var flag bool = true

	for i:=0; i<len(f)-2; i++ {
		if (string(f[i+1]) == "x" || string(f[i+1]) == "=") && flag{
			num, _ = strconv.Atoi(f[count:i+1])
			flag = false
		}else if unicode.IsNumber(rune(f[i+1])) && string(f[i]) == "+" && i>0{
			flag = true
			count = i+1
		}

		if a == 0 {
			a, num = num, 0
		} else if b == 0 {
			b, num = num, 0
		} else {
			c, num = num, 0
		}
	}

	var delta float64 = float64(b*b) - float64(4*a*c)
	var x1, x2 float64

	if delta > 0 {
		x1 = (-float64(b) + math.Sqrt(delta))/(2*float64(a))
		x2 = (-float64(b) - math.Sqrt(delta))/(2*float64(a))

		fmt.Printf("Le soluzioni reali sono due distinte e sono: %.3f, %.3f\n", x1, x2)
		
		if x1 > (soglia+epsilon){
			fmt.Println("La prima soluzione è maggiore della soglia + epsilon")
		}
		if x2 > (soglia+epsilon){
			fmt.Println("La seconda soluzione è maggiore della soglia + epsilon")
		}
	} else if delta == 0 {
		x1 = (-float64(b) + 0)/(2*float64(a))

		fmt.Println("Le soluzioni reali sono due coincidenti ed è: %.3f\n", x1)

		if x1 > (soglia+epsilon){
			fmt.Println("La soluzione è maggiore della soglia inposta")
		}
	} else if delta < 0{
		fmt.Println("Non esistono soluzioni reali")
	}
}