package main

import (
	"fmt"

	"github.com/JuanRodriguez84/godesde0/variables"
)

func main() {
	fmt.Println("func main")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	variables.ConviertoaTexto(1588)
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println("main - estado = ", estado)
	fmt.Println("main - texto = ", texto)
}
