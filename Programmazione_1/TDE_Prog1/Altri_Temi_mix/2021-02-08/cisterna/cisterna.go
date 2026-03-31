package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Cisterna struct {
	livello, x, y, z float64
}

//in questa funzione passo come primo argomento una variabile di tipo puntatore alla struct Cisterna
// facendo questo, nel programma potrò quindi modificare il valore di una variabile di tipo Cisterna già dichiarata in precedenza

func variazione(cisterna *Cisterna, lt int) error {

	var err error

	volumeCisternaD := ((cisterna.x * cisterna.y * cisterna.z) / 1000) - ((cisterna.livello * cisterna.x * cisterna.y) / 1000) //L

	volumeCisterna := (cisterna.x * cisterna.y * cisterna.livello) / 1000

	if lt > 0 {

		if float64(lt) <= volumeCisternaD {

			volumeCisterna := (volumeCisternaD - float64(lt)) //L

			cisterna.livello = cisterna.z - (volumeCisterna*1000)/(cisterna.x*cisterna.y)

			return err

		} else {

			stringa := fmt.Sprintf("puoi mettere max " + strconv.Itoa(int(volumeCisternaD)) + " litri")

			// questa funzione del pacchetto errors serve semplicemente a stampare un errore in una variabile di tipo stringa
			// se non è strettamente richiesto dal testo dell'esercizio puoi anche usare una semplice fmt.Println

			return errors.New(stringa)
		}

	}

	// la funzione math.Abs serve a fare il valore assoluto di un'operazione

	if math.Abs(float64(lt)) <= volumeCisterna {

		volumeCisternaBoh := (volumeCisterna + float64(lt)) //L

		cisterna.livello = (volumeCisternaBoh * 1000) / (cisterna.x * cisterna.y)

	} else {

		stringa := fmt.Sprintf("puoi prendere max " + strconv.Itoa(int(volumeCisterna)) + " litri")

		return errors.New(stringa)

	}

	return err

}

func contenuto(cisterna Cisterna) (litri int) {

	return int((cisterna.livello * cisterna.x * cisterna.y) / 1000)
}

// questa funzione si chiama metodo, e spesso i metodi sono presenti nei temi d'esame. a differenza delle funzioni normali
// per i metodi bisogna scrivere dopo la parola func, tra parentesi tonde, il nome di una variabile e il suo tipo, poi
// il nome del metodo, in questo caso String(). quando si dovrà chiamare il metodo, si scriverà il nome di una variabile a cui si
//vuole applicare il metodo, un punto, e il nome del metodo. In questo caso quindi scriverò cisterna.String()

func (cisterna Cisterna) String() {

	fmt.Printf("cisterna %.2f cm x %.2f cm x %.2f cm\nlivello attuale : %.2f cm, litri %d\n", cisterna.x, cisterna.y, cisterna.z, cisterna.livello, int((cisterna.x*cisterna.y*cisterna.livello)/1000))
}

func main() {

	x, err := strconv.ParseFloat(os.Args[1], 64)
	y, err1 := strconv.ParseFloat(os.Args[2], 64)
	z, err2 := strconv.ParseFloat(os.Args[3], 64)
	var interi int
	cist := Cisterna{0, x, y, z}

	if x < 0 || y < 0 || z < 0 {

		fmt.Println("tutti gli argomenti devono essere > 0")
		return

	}

	if err != nil || err1 != nil || err2 != nil { //ovvero sfrutto la variabile errore che mi ritorna ParseFloat per vedere se ci sono numeri che non sono float

		fmt.Println("tutti gli argomenti devono essere numerici")
		return
	}

	cist.String() // utilizzo il metodo sulla variabile cist di tipo Cisterna

	// entro in un ciclo perchè devo leggere da standard input una serie di valori che non so quanto è lunga

	for {
		fmt.Scan(&interi)

		// esco dal ciclo di lettura quando il programma legge 9999

		if interi == 9999 {

			break
		}

		// nella funzione variazione bisogna passare come primo parametro un puntatore.const
		// facendo tutti i passaggi quindi nel programma bisognerebbe dichiarare una variabile ad esempio p *Cisterna
		// assegnare a questa variabile l'indirizzo in memoria della variabile a cui voglio riferirmi : quindi scrivere p = &cist
		//in questo caso quindi ho messo che p punta all'indirizzo in memoria della variabile cist di tipo Cisterna
		//infine passare p come primo parametro della funzione variazione
		//al posto di fare tutti questi passaggi, è anche possibile passare come primo parametro della funzione direttamente l'indirizzo
		//della variabile a cui vogliasmo riferirci, quindi in questo caso scrivo direttamente &cist, che sarebbe p.

		err := (variazione(&cist, interi))

		if err != nil {

			fmt.Println(variazione(&cist, interi))
		} else {

			cist.String() // uso di nuovo il metodo
		}
	}

}

/*
Cisterna
========


Scrivere un programma 'cisterna.go' dotato di:

- una struttura 'Cisterna' con 3 campi float64 per le tre dimensioni (in cm) (NOTA SERVIRA' SPECIFICARE NOME E ORDINE DEI CAMPI?) e un ulteriore campo float64 che indica il livello (in cm) a cui arriva l'acqua nella cisterna

- una funzione

    variazione(cisterna *Cisterna, lt int) error

  che, se lt è positivo e c'è abbastanza spazio nella cisterna, versa lt litri (***) nella cisterna e restituisce errore nil, altrimenti non versa niente nella cisterna e restituisce un errore (vedi esempio) segnalando la capienza in litri ancora disponibile; se invece lt è negativo, se ci sono abbastanza litri, prende lt litri dalla cisterna e restituisce errore nil, altrimenti non prende niente dalla cisterna e restituisce un errore (vedi esempio) segnalando la disponibilità (in litri).

- una funzione
	func contenuto(cisterna Cisterna) (litri int)
  che restituisce il numero di litri contenuti nella cisterna

- un metodo String() per Cisterna che restituisce una descrizione della cisterna (vedi esempio sotto)

- una funzione main() che legge da linea di comando tre numeri (float64) che rappresantano larghezza, lunghezza e altezza (in quest'ordine) di una cisterna a base rettangolare.
Se i dati sono tutti > 0, crea una cisterna vuota (livello = 0) con quelle dimensioni e la stampa (vedi es.); altrimenti segnala l'errore "tutti gli argomenti devono essere >0" (se anche uno solo è minore o uguale a 0) o "tutti gli argomenti devono essere numerici" (se anche uno solo non è numerico) e termina.
Poi legge da standard input una sequenza di interi, che rappresentano litri, terminata da 9999: quando il numero letto è positivo, viene messo nella cisterna un numero di litri corrispondenti, viceversa quando il numero è negativo, i litri vengono tolti, sempre che ciò sia possibile, altrimenti viene segnalato un errore (senza però terminare).
Dopo ogni variazione effettiva,  il programma stampa lo stato della cisterna (vedi es.).
Se invece la variazione della cisterna non va a buon fine, il programma non ristampa il suo stato,

*** Nota: un litro corrisponde a un decimetro cubo

Suggerimenti:
1. per creare un errore, utilizzare la funzione New del package 'errors' (vedere la documentazione relativa).

Esempio
=======

(nota: per distinguere input da output, l'input è preceduto da >>>, che però non andrà digitato)

$ ./cisterna 50 100 150
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 0.00 cm, litri 0
>>>1000
puoi mettere max 750 litri
>>>500
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 100.00 cm, litri 500
>>>-600
puoi prendere max 500 litri
>>>-300
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 40.00 cm, litri 200
>>>9999
*/
