package ejercicio_interfaces

import (
	"fmt"

	"github.com/JuanRodriguez84/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) { // Hombre y Mujer automaticamente son Humano ya que implementan los metodos definidos en Humano
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando\n", hu.Sexo())
}

func SerVivo(sv interfaces.SerVivo) { // Hombre y Mujer automaticamente son Humano ya que implementan los metodos definidos en Humano
	sv.EstaVivo()
	fmt.Println(sv.EstaVivo())
}
