package funciones

import (
	"fmt"
)

func Calculos() {

	//	var numero1 int
	//	var numero2 int

	// se puede hacer en una linea

	// var numero1, numero2 int

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	// para llamar a la función es:

	fmt.Println(suma(10, 25))

}

func Calculos2() {

	var numero3 int = 32
	var numero4 int = 243

	// se crea para tener un codigo acotado y no tener muchas funciones

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4 // si se crea una funcion por fuera las variables numero3 y numero4 no se verian por el scope
	}

	fmt.Println(calculo(10, 25))

	// no se puede cambiar la firma de la fucnion, ejemplo :    calculo := func(numero1 int, numero2 int, numero3 int) int { ...

	// al rasignar el calculo es con = , no con :=
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 // si se crea una funcion por fuera las variables numero3 y numero4 no se verian por el scope
	}

	// para llamar a la función es:

	fmt.Println(calculo(10, 25))

}

// func   devolviendo una funcion que devuelve int   :

func Calculos3() func(int, int) int {

	var numero3 int = 32
	var numero4 int = 243

	// se crea para tener un codigo acotado y no tener muchas funciones

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4 // si se crea una funcion por fuera las variables numero3 y numero4 no se verian por el scope
	}

	fmt.Println(calculo(10, 25))

	// no se puede cambiar la firma de la fucnion, ejemplo :    calculo := func(numero1 int, numero2 int, numero3 int) int { ...

	// al rasignar el calculo es con = , no con :=
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 // si se crea una funcion por fuera las variables numero3 y numero4 no se verian por el scope
	}

	// para llamar a la función es:

	fmt.Println(calculo(10, 25))

	return calculo

}
