package main

import (
	"fmt"

	"github.com/JuanRodriguez84/godesde0/teclado"
)

func main() {
	fmt.Println("func main")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/* variables.ConviertoaTexto(1588)
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println("main - estado = ", estado)
	fmt.Println("main - texto = ", texto)

	os := runtime.GOOS // GOOS, dice que el sistema operativo de el programa que está corriendo lo tiene esta propiedad, con esto se sabe si es windows o linux
	fmt.Println("main - os = ", os)

	if os == "Linux." { // == comparador   si es in solo = es asignación
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
	}

	// comparadores ==
	//              !=  // diferente
	//              >=
	//              <=
	//              <=
	//              >=
	//              &&  // And
	//              ||  // OR

	// se puede hacer asi
	//   ASIGNACION       ;  PREGUNTA
	if os2 := runtime.GOOS; os2 == "Linux." || os2 == "OS X." { // == comparador   si es in solo = es asignación
		fmt.Println("Esto no es windows, es ", os2)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os3 := runtime.GOOS; os3 { // se puede hacer la asignación dentro al igual que el if
	case "Linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os3) // Printf permite formatear el texte de alguna manera. %s   significa que el argumento que le voy a pasar es un string
	} */

	/* 	numero, texto := ejercicios.ConversorNumerico("001")
	   	fmt.Println("numero", numero)
	   	fmt.Println("texto", texto) */

	teclado.IngresoNumeros()

}
