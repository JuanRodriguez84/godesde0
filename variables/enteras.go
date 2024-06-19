package variables

import "fmt"

func MuestroEnteros() { // metodos no tienen argumentos de retorno
	var intcomun int
	intde32 := int32(10)

	intde64 := int64(100)

	// intde64 = intcomun En go int, int32 y int64 a pesar de ser integer son tipos de datos diferemtes y no se puede hacer algo como eso

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
