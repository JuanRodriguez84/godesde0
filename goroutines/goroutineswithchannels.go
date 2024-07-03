package goroutines

import (
	"fmt"
	"strings" // paquete para todo lo que tiene que ver con funciones de strings, por un lado está el tipo de dato string y por otro lado la libreria strings que esta llena de metodos dedicados a los strings
	"time"
)

func MiNombreLentoWithChannel(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") // vector llamado letras

	for _, letra := range letras { // _  no interesa el primer caracter del range  , // el range se usa siempre con el for, es una especia de foreach
		time.Sleep(1000 * time.Millisecond) // un segundo
		fmt.Println(letra)
	}

	canal1 <- true // la forma de asignarle un valor a un channel es con <-    // canal1 <- true   cuando no se aigna a ninguna variable
	// cuando le asignamos un valor explicitamente va a disparar una especie de trigger para que desde donde hemos llamado a
	// goroutines sepa que el valor tuvo un valor nuevo , puede ser bool, int , string lo importante es establecer un nuevo valor
	//
}
