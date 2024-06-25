package arreglos_slices

import (
	"fmt"
)

//var tabla []int // de esta manera se defince un slices, uja de las diferencias entre vector y un slices, es que el slice no tiene una cantidad de dimensión

var tabla [10]int // vector de 10 elementos, comienzan en cero e iria de la 0 a la 9
func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	fmt.Println(tabla)
}

var tabla2 [10]int = [10]int{10, 0, 59}

// asignar valores al array de manera directa
func MuestroArreglos2() {
	tabla2[7] = 33
	tabla2[2] = 54

	fmt.Println(tabla2)
}

func MuestroArreglosVariante() {
	tabla[7] = 33
	tabla[2] = 54
	tabla3 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla3)
}

// la manera de recorrer vectores
func RecorrerArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	tabla3 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla3)
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

}

// como crear una matriz que es un vector de vectores
var matriz [20][30]int

func Matriz() {
	matriz[7][24] = 15
	fmt.Println(matriz)
}
