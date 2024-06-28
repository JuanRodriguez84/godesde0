package main

import (
	"fmt"

	"github.com/JuanRodriguez84/godesde0/defer_panic_recover"
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

	/* teclado.IngresoNumeros() */

	// puede ir sin terminar nunca de esta manera:

	//fmt.Println(tablaparamanejodearchivos.TablaDeMultiplicarConRetorno())

	// files.GrabaTablaEnArchivo()

	// files.SumaTablaEnArchivo()
	// files.LeoArchivo2daForma()

	/* 	// funciones anonimas

	   	funciones.Calculos()

	   	fn := funciones.Calculos3()
	   	resultado := fn(5, 7)
	   	fmt.Println("Resultado:", resultado)

	   	// se puede llamar directamente :
	   	fmt.Println("****Directamente ****")
	   	funciones.Calculos3()

	   	// Closures

	   	fmt.Println("****Closures ****")
	   	funciones.LlamarClosure() */

	/* 	// Recursion

	   	funciones.Exponencia(2) */

	/* // Arreglos estáticos

	fmt.Println("****Arreglos estaticos****")

	arreglos_slices.MuestroArreglos()

	arreglos_slices.MuestroArreglos2()

	arreglos_slices.MuestroArreglosVariante()

	arreglos_slices.RecorrerArreglos()

	arreglos_slices.Matriz()

	//Slices

	fmt.Println("****Slices****")

	arreglos_slices.MuestroSlice()
	arreglos_slices.MuestroSlice2()
	arreglos_slices.Capacidad()
	arreglos_slices.Capacidad2() */

	/* 	fmt.Println("****Mapas****")

	mapas.MostrarMapas() */

	/* 	fmt.Println("****Interfaces****")

	   	Juan := new(modelos_estructuras.Hombre)
	   	ejercicio_interfaces.HumanosRespirando(Juan)
	   	// una interface puede heredar de otra interface.
	   	ejercicio_interfaces.SerVivo(Juan)

	   	Eva := new(modelos_estructuras.Mujer)
	   	ejercicio_interfaces.HumanosRespirando(Eva)
	   	// una interface puede heredar de otra interface.
	   	ejercicio_interfaces.SerVivo(Eva) */

	fmt.Println("****DEFER - PANIC & RECOVER****")

	defer_panic_recover.VemosDefer()
	// defer_panic_recover.EjemploPanic() // se comentara para que funcione  EjemploRecover  ya que este metodo tambien tiene panic y no deja avanzar
	defer_panic_recover.EjemploRecover()

}
