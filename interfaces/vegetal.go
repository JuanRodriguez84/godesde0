package interfaces

type Vegetal interface { // interface palabra reservada
// aca van las definiciones de funciones , que es lo que hace un vegetal

	ClasificaciónVegetal() string
	EstaVivo()bool  // de la interface SerVivo
}