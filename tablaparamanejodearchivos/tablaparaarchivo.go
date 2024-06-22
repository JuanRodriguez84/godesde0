package tablaparamanejodearchivos

import (
	"bufio" // librerio io (input output)  tiene que ver con lectura no solo por teclado si no que tambien hace tratamiento de archivos y demas, en este caso se usa para teclado
	"fmt"   // viene en el statdard de GO no hay que importarlo
	"os"    // paquete del sitema operativo
	"strconv"
)

func TablaDeMultiplicarConRetorno() string { // es un unico dato que devuelve asi que no hace falta colocar parentesis

	scanner := bufio.NewScanner(os.Stdin)

	var numero1 int
	var err error
	var texto string

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

	for i := 0; i <= 10; i++ {
		//	fmt.Printf("%d X %d = %d \n", numero1, i, numero1*i)  no se utiliza por que devuelve un entero y se necesita un string
		//fmt.Sprintf("%d X %d = %d \n", numero1, i, numero1*i) // Sprintf  devuelve un string
		//fmt.Fprintln()   // guarda en un archivo el resultado

		// concatenar +=

		texto += fmt.Sprintf("%d X %d = %d \n", numero1, i, numero1*i)

	}
	return texto

}
