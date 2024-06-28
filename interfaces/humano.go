package interfaces

type Humano interface { // interface palabra reservada
// aca van las definiciones de funciones , que es lo que hace un humano

	Respirar()
	Pensar()
	Comer()
	Sexo() string
	EstaVivo()bool // de la interface SerVivo
}