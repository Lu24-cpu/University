/*
# TESTO

Realizzare un programma Go (nome file 'tail.go') che implementi un semplice 'tail', comando Unix che estrae le ultime N righe di un file di testo.

Il programma deve prendere due parametri da linea di comando (in quest'ordine):
- numero N di linee da estrarre
- nome del file da elaborare

e stampare su standard output le ultime N righe del file. Se il file ha meno di N righe, il programma stamperà tutte quelle che ci sono.

Nota bene: NON va implementato invocando da Go il comando 'tail' di sistema.

## Esempio esecuzione

(presuppone il 'tail.go' già compilato, 'uno.input' è presente in questa directory)

$ ./tail 3 uno.input
remaining essentially unchanged. It was popularised in the 1960s with the release
of Letraset sheets containing Lorem Ipsum passages, and more
recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

*/

package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)
func main(){
	/*
	if len(os.Args) < 3 {
		return
	}
	*/
	var dim_file int
	n, err := strconv.Atoi(os.Args[1])
	if err != nil{
		return
	}

	fp, err := os.Open(os.Args[2])
	if err != nil{
		return
	}
	
	scanner := bufio.NewScanner(fp)
	for scanner.Scan(){ dim_file++ }
	fp.Close()

	fp, _ = os.Open(os.Args[2])
	scanner = bufio.NewScanner(fp)
	i := 0
	for scanner.Scan(){
		if i >= (dim_file - n){
			fmt.Println(scanner.Text())
		}
		i++
	}
	fp.Close()
}