package mapas

import (
	"fmt"
)

func MostrarMapas() {

	// un mapa se peude definir de dos maneras
	// 1. con las instrucción make

	paises := make(map[string]string) // palabra map , es palabra reservada en go para hablar de mapas, en este ejemplo [string] es el tipo de dato de la llave y string es el tipo de dato del valor
	fmt.Println(paises)

	// asignar valores
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	// se puede colocar una longitud a un mapa :
	// paises2 := make(map[string]string, 5)  // va a tener un tamaño de 5

	// 2. la otramanera de crear un mapa es:

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca":        07,
		"Z":           07,
		"A":           05,
		"B":           04,
		"x":           03,
	}

	// los muestra en orden alfabetico la clave

	fmt.Println(campeonato)

	// recorrer la coleccion, aca no respeta orden ni de clave ni de valor para imprimir
	for equipo, puntaje := range campeonato { // el range se usa siempre con el for, es una especia de foreach
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje) // %s para procesar strings ,  %d  significa que es un numerico
	}

	// eliminar un elemento de un mapa con la función delete
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	// cuando se busca por clave que devuelve?
	puntuacion, existe := campeonato["Juventus"] // la consulta del mapa devuleve en realidad dos caracteristicas - 1.valor que tiene la clave y un bolenao que indica si existe esa clave o no

	fmt.Printf("El puntaje capturado es %d, y el equipo existe es = %t \n", puntuacion, existe) // %t  ->  para armar un mensaje con variable booleana

	puntuacion, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe es = %t \n", puntuacion, existe)
}
