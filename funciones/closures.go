package funciones

import (
	"fmt"
)

func Tabla(valor int) func() int {
	numero := valor
	secuencia:=0

	return func() int {
		secuencia++
		return numero*secuencia
	}
}

// función que devuleve una fucnión
func LlamarClosure(){
	tabladel := 2
	Mitabla := Tabla(tabladel)

	for i:=1; i<=10; i++{
		fmt.Println(Mitabla())
	}
}