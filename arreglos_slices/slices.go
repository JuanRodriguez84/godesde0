package arreglos_slices

import (
	"fmt"
)

// slice son vectores a los cuales no se les crea una capacidad inicial
var tablaS []int = []int{2, 5, 4}

func MuestroSlice() {

	fmt.Println(tablaS) // cuando se imprime no llena de ceros ya que el slice tiene 3 elementos, ajusta la dimensión del arreglo por la asignación
}

// como google maneja el performance, por que asignación y redefinición de arreglos en memoria para los lenjuages es muy costoso a nivel de performance
// cuando hay mucha recurrencia este tipo de operaciones se siente

// por eso se va a ver como se maneja los elementos dinamicos, por que en definitiva un slice es un arreglo dinamico

var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice2() {

	// Go permite crear un slice tomando un arreglo como modelo y tambien permite tomar algunos elementos del arreglo para crear un slice

	// Slice desde un arreglo, se dice que es slice por que uno de los datos no se está pasando, no se pasa hasta donde se va a asignar ni desde donde se va a asignarm, ejemplo:
	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posicion 3 hasta el final ---  3:  ->  significa desde el elemento 3 hasta el final
	porcion2 := arreglo[:5]  // Slice creado desde un vector, desde la posicion 0 hasta la posición 5
	porcion3 := arreglo[6:7] // Slice creado desde un vector, desde la posicion 6 hasta la posición 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

// los slice manejan dos valores, la dimensión , osea la cantidad de elementos y un valor que no es visible que se llama capacidad
// la capacidad, es el espacio en memoria que go reserva para crear ese slice que no es la cantidad de elementos y go no lo hacer por ahorro de
// memoria, go lo hace por performance
func Capacidad() {

	// los slice a diferencia de los arreglos son dinamicos, lo que significa que puedo adicionarle mas capacidad de la definida a cuando se declaró
	// los arreglos en tiempo de ejecución no se pueden modificar la capacidad  ya que se inicialican con una cantidad especifica, ejemplo : var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}
	// mientras que los slice si se pueden agrandar conforme a necesidades en tiempo de ejecución

	// la instrucción "make"  permite crear un slice vacio  y darle una capacidad, ya que cuando se crea como variable no se tiene manera de indicarle la capacidad
	// make sireve pero todo tipo de estructuras
	elementos := make([]int, 5, 20)                                         // slice de 5 elementos y 20 de capacidad, el tercer parametro es la capacidad que va a reservar en memoria
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos)) // cap es una fucnión en GO para obtener la capacidad de un slices

}

// la manera de agregar elementos a un slice:
func Capacidad2() {
	nums := make([]int, 0, 0)

	// para agregar es con append

	for i := 0; i < 100; i++ {
		nums = append(nums, i) // en esta parte se adiciona
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))

	// el resultado es Largo 100, Capacidad 128, a medida que va creciendo va reservando. (128) es el resultado de cómo Go gestiona el crecimiento (de manera exponencial) del slice en ese contexto específico.
	// si el largo fuera de 200 la capacidad seria de 256, al parecer despues del multiplo de 8 le suma 8, es decir, cuando esta en 9 (8+1) el largo la capacidad sube a 16, cuando esta en 17 (16+1) a 32, cuando esta en 33 (32+1) a 64, cuando esta en 65 (64+1) a 128. etc
	// ese calculo es el que realiza cuando hay una sobrecarga importante de append cuando crece el slice y hace que reserve más capacidad,
	// Go hace esto por que para el lenguaje es mas costoso ir a reservar memoria que utilizarla, es mas costoso cambiar la dirección de memoria de una variable por que el espacio que
	// reservo no alcanza, esta operación es de muy bajo nivel

	// los slice en tipo de ejecución puedo jugar con sus elemento, ampliarlo, reducirlo
}
