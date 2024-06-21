package ejercicios

import (
	"bufio" // librerio io (input output)  tiene que ver con lectura no solo por teclado si no que tambien hace tratamiento de archivos y demas, en este caso se usa para teclado
	"fmt"   // viene en el statdard de GO no hay que importarlo
	"os"    // paquete del sitema operativo
	"strconv"
)

func TablaDeMultiplicar() {

	var numero1 int
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	/*	var isNumber bool = false

			//  FORMA 1

		 	for !isNumber {
				fmt.Println("Ingrese Numero de Tabla : ")
				if scanner.Scan() { // si el usuario ingreso algo por teclado
					numero1, err = strconv.Atoi(scanner.Text()) // lo que entra en bufio entra en modo texto, por eso se convierte a entero
					if err == nil {
						isNumber = true
					} else {
						fmt.Println("Error, digite nuevamente")
					}
				}
			} */

	// FORMA 2

	for {
		fmt.Println("Ingrese Numero de Tabla : ")
		if scanner.Scan() { // si el usuario ingreso algo por teclado
			numero1, err = strconv.Atoi(scanner.Text()) // lo que entra en bufio entra en modo texto, por eso se convierte a entero
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	fmt.Println("Tabla del ", numero1)

	for i := 0; i <= 10; i++ {
		//	fmt.Println(numero1, " X ", i, " = :", numero1*i)
		fmt.Printf("%d X %d = %d \n", numero1, i, numero1*i) // %d indica que va a recibir un numero, de tipo base 10 por eso la d
	}

}
