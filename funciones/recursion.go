package funciones

import (
	"fmt"
)

// recursión es una funcion que se llama a si misma N veces

func Exponencia(valor int){

	// cuando se trabaja con recursión lo primero que se tiene que trabajar es la salida de la recursión
	if valor > 10000000 {
		return
	}

	fmt.Println(valor)

	Exponencia(valor * 2)
}