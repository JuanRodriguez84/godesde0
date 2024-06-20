package ejercicios

import (
	"strconv" // viene en el statdard de GO no hay que importarlo
)

func ConversorNumerico(texto string) (int, string){
	//num, _ := strconv.Atoi(texto)    // num, _  :  el argumento que me esta pasando no se asigna a ninguna variable y strconv.Atoi convierte de alfabetico a integer
	//if num > 100{
	//	return num, "Es mayor a 100"
	//} else{
	//	return num, "Es menor a 100"
	//}

	// para manejar el error sería :

	num, err := strconv.Atoi(texto)    // num, _  :  el argumento que me esta pasando no se asigna a ninguna variable y strconv.Atoi convierte de alfabetico a integer
	
	if err!= nil{
		return 0, "hubo un error" +  err.Error()   // err.Error()  devuelve el mensaje de error
	}
	
	if num > 100{
		return num, "Es mayor a 100"
	} else{
		return num, "Es menor a 100"
	}
}
