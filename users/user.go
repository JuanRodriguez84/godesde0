package users

import (
	"time" // dato de tipo fecha
	"fmt"
	"github.com/JuanRodriguez84/godesde0/modelos_estructuras"
)

func AltaUsuario(){

	// En programación orientada a objetos, las definiciones no son objetos hasta que uno no los cree con un new, en go pasa lo mismo
	u := new(modelos_estructuras.User)  // asi se crea un objeto
	u.AddUser(10, "juan", time.Now(), true)
	// aca pueden colocarse logicas de validaciones, ejemplo conectarse a una BD para validar si el usuario ya está en BD
	fmt.Println(u)  // imprime algo como esto &{10 juan 2024-06-28 08:46:50.2758035 -0500 -05 m=+0.004014201 true} , donde  &  indica que es una estructura
}