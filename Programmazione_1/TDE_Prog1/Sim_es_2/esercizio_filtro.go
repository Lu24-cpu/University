package main

import (
	"fmt"
	"os"
)

func main(){
	var k int
	stringa := os.Args[1]
	
	l := len(stringa)-1

	for i :=0; i<l/2; i++ {		
		for k=0; k<l/2-i; k++{
			fmt.Print(" ")
		}
		fmt.Println(stringa[:(l/2)-k+i+1])
	}
	
	fmt.Println(stringa)

	var count int =1
	for i := l/2; i<l; i++ {		
		for k=i+1; k>l/2; k--{
			fmt.Print(" ")
		}
		fmt.Println(stringa[:l-i+k-count])
		count++
	}
}
