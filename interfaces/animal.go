package interfaces

type Animal interface { // interface palabra reservada
// aca van las definiciones de funciones , que es lo que hace un animal

	Respirar()
	Comer()
	EsCarnivoro() bool
	EstaVivo()bool  // de la interface SerVivo
}