package teclado

import (
	"bufio" // librerio io (input output)  tiene que ver con lectura no solo por teclado si no que tambien hace tratamiento de archivos y demas, en este caso se usa para teclado
	"fmt"   // viene en el statdard de GO no hay que importarlo
	"os"    // paquete del sitema operativo
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error // tipo de dato error

func IngresoNumeros() {

	scanner := bufio.NewScanner(os.Stdin) // captura de teclado , el argumento que pide es de donde leer los datos, por que lo puede leer de archivos
	// lo puede leer de dispositivos ipo colectores de datos, lectores de codigo de barras, por eso se utiliza os, el paquete
	// os  con el estandard input (Stdin), el estandard input en cualquier sistema, o lenguaje es el teclado el Stdout es la pantalla
	// o la consula
	// apartir de este momento scanner puede leer todas las veces que yo quiera

	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() { // si el usuario ingreso algo por teclado
		numero1, err = strconv.Atoi(scanner.Text()) // lo que entra en bufio entra en modo texto, por eso se convierte a entero
		if err != nil {
			//fmt.Printf("Dato ingresado incorrecto")
			// o si se desea abortar la aplicación se hace asi:
			panic("Dato ingresado incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() { // si el usuario ingreso algo por teclado
		numero2, err = strconv.Atoi(scanner.Text()) // lo que entra en bufio entra en modo texto, por eso se convierte a entero
		if err != nil {
			panic("Dato ingresado incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() { // si el usuario ingreso algo por teclado
		leyenda = scanner.Text() // lo que entra en bufio entra en modo text
	}

	fmt.Println(leyenda, numero1*numero2)

}
