package main

import "fmt"

func Fattoriale(n int) int {
	if n == 1 {
		return 1
	}

	return (n * Fattoriale(n-1))
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return (Fibonacci(n-1) + Fibonacci(n-2))
}

func MulconSum(n, m int) int {
	if m == 1 {
		return n
	}

	return (n + MulconSum(n, m-1))
}

func Somma(valori []int) int {
	if len(valori) < 3 {
		return 0
	}

	sum := valori[0] + valori[1] + valori[2]

	return (sum + Somma(valori[1:]))
}

func main() {
	var n, m int
	var valori []int = []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Inserire un valore per il fattoriale e per la serie di Fibonacci")
	fmt.Scan(&n)

	fmt.Println("Ecco il fattoriale:", Fattoriale(n), "e il valore della serie di Fibonacci:", Fibonacci(n))

	fmt.Println("Ora una moltiplicazione fatta con somme, inserisci i due addendi:")
	fmt.Scan(&n, &m)

	fmt.Println("Moltiplicazione:", MulconSum(n, m))

	fmt.Println("Ora ti faccio vedere l'esercizio del mio orale di programmazione.\nDato lo slice:", valori, "Calcolare la somma di 3 elementi in 3 e poi fare la complessiva di tutte")

	fmt.Println("Somma:", Somma(valori))
}
