package variables

import (
	"fmt"

	"strconv" // viene en el statdard de GO no hay que importarlo
	"time"    // para trabajar con fechas y horas se utiliza este paquete, time  es un paquete que viene dentro de GO
)

//var nombre string     // NO se ve la variable en todo el paquete por que comienza en mayuscula

var Nombre string // se ve la variable en todo el paquete por que comienza en mayuscula
var Estado bool
var Sueldo float32 // dependiendo de la cantidad de digitos y decimales elegimos entre float64 o float32
var Fecha time.Time

func RestoVariables() {
	// MuestroEnteros() // es visible en todo el paquete variables el metodo MuestroEnteros

	Nombre = "pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado) // en Println el arguemnto es any por eso permite cualquier tipo de dato, ya sea boolean o float
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {
	// var texto string               // se puede asignar directamente
	// texto = string(numero)         // Imprime caracter extraño
	texto := strconv.Itoa(numero) // funcion strconv.Itoa para Convertir de integer a alfanumerico
	fmt.Println("ConviertoaTexto - texto = ", texto)
	return true, texto
}
