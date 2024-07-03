package goroutines

import (
	"fmt"
	"strings" // paquete para todo lo que tiene que ver con funciones de strings, por un lado está el tipo de dato string y por otro lado la libreria strings que esta llena de metodos dedicados a los strings
	"time"
)

func MiNombreLento(nombre string){
	letras := strings.Split(nombre, "")       // vector llamado letras

	for _, letra := range letras { // _  no interesa el primer caracter del range  , // el range se usa siempre con el for, es una especia de foreach
		time.Sleep(1000* time.Millisecond)  // un segundo
		fmt.Println(letra)
	}	
}