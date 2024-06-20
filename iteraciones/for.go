package iteraciones

import (
	"fmt"
)

func ForInfinito() {

	for {
		fmt.Println("hola en for")
	}

}

func ForConBreak() {
	// puede terminar de esta manera:

	for {
		fmt.Println("hola en for")
		break // terminar con for
	}
}

func IterarForConIncrementos() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++ // incrementar i en uno (1)
	}
}

func IterarDeManeraSencillaFor() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func IterarConSaltosGrandesEnFor() {

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}
}

func IterarDeParaAtras() { // empezando desde 10

	for i := 10; i > 1; i-- {
		fmt.Println(i)
	}
}

func IterarConSaltosGrandesDeParaAtras() { // empezando desde 100

	for i := 100; i > 10; i -= 5 {
		fmt.Println(i)
	}
}

func IterarConBreak() {

	for i := 100; i > 10; i -= 5 {
		if i == 80 {
			break
		}
		fmt.Println(i)
	}
}


func IterarConContinue() {

	for i := 100; i > 10; i -= 5 {
		if i == 80 {
			continue      // no continua y pasa a la siguiente secuencia del for, no mostraria el 80
		}
		fmt.Println(i)
	}
}
