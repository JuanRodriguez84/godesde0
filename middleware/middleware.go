package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	//result := operacionesMidd(sumar)(2, 3) -> en Go utiliza funciones anónimas y es un ejemplo de cómo se puede implementar la técnica conocida como "middleware" en Go.
	result := operacionesMidd(sumar)(2, 3) // llamar la funcion sumar con los argumentos 2, 3
	fmt.Println(result)
	result = operacionesMidd(restar)(10, 6) // llamar la funcion restar con los argumentos 10, 6
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 4) // llamar la funcion multiplicar con los argumentos 2, 4
	fmt.Println(result)
}

func operacionesMidd(funcionMia func(int, int) int) func(int, int) int { // Los middleware reciben un tipo de dato determinado y deben devolver los mismos
	// un middleware es un pasamanos donde se recibe y asi como se recibe se hace algo pero se debe devolver lo mismo,
	// para este ejemplo no se puede devolver un boolean para que en algun momento se ejecute sumar, restar o multiplicar
	return func(a, b int) int { // son los argumentos de la funcion sumar, restar y multiplicar
		// aca se puede preguntar por a o por b, se puede validar si "a" es mayor a cierto numero , mostrar un mensaje por consula
		fmt.Println("Inicio de operacion") // se devuleve lo mismo como si no hubiera echo nada, lo unico que se hace en este ejemplo es el Println
		return funcionMia(a, b)
	}
}
