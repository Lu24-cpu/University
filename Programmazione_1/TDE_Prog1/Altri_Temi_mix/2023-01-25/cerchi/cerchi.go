package main

import(
	"fmt"
	"math"
	"bufio"
	"os"
)

type Cerchio struct{
	nome string
	x,y,raggio float64
}

func newCircle(descr string) Cerchio{
	// nome x y raggio
	var x, y, raggio float64
	var nome string
	fmt.Sscan(descr, &nome, &x, &y, &raggio)
	return Cerchio{nome, x, y, raggio}
}

func distanzaPunti(x1,y1,x2,y2 float64) float64{
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func maggiore(cerc1, cerc2 Cerchio) bool{
	return cerc1.raggio > cerc2.raggio
}

func tocca(cerc1, cerc2 Cerchio) bool{
	distanza := distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)
	
	// check se sono esterni: distanza = somma dei raggi
	somma := cerc1.raggio + cerc2.raggio
	if distanza < somma + MIN_D && distanza > somma - MIN_D {
		return true
	}
	// check se sono interni: distanza = differenza dei raggi
	differenza := _abs(cerc1.raggio - cerc2.raggio)
	if distanza < differenza + MIN_D && distanza > differenza - MIN_D {
		return true
	}
	return false	
}


func _abs(n float64) float64{
	if n < 0 {
		return -n
	}
	return n
}

const MIN_D = 1e-6

func main(){
	var past, current Cerchio
	var input string
	var tangente, grande bool

	past = Cerchio {"", 0.0, 0.0, 0.0}
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		input = scanner.Text()
		current = newCircle(input)
		fmt.Print(current)
		
		tangente = tocca(current, past)
		if tangente{
			fmt.Print(" tangente,")
		}else{	
			fmt.Print(" non tangente,")
		}
		grande = maggiore(current, past)
		if grande {
			fmt.Print(" maggiore")
		}else{	
			fmt.Print(" minore o uguale")
		}
		fmt.Println()
		past = current
	}
	
}