package modelos_estructuras

import (
	"time" // dato de tipo fecha
)

// estructura de las estructuras, sintaxis de las estructuras es iniciando con la palabra type
// el paquete modelos_estructuras se va a utilizar para colocar definiciones y estructuras, es un standard , en los desarrollos de GO se hacen asi
// donde se tinen todas las estructuras en un solo lugar  y la logica en otro paquete

type User struct { // palabra reservada struct
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// si como en POO se quiere que la estructura o la clase tenga funcionalidad, ejemplo AddUser es con
// func pero entre el parentesis se indica la estructura que se está referenciando


/* func (this User) AddUser(id int, name string, createdAt time.Time, status bool) { // con this se refiere a User, no es obligatorio la palabra this se usa por convensión con otros lenguajes, this o self en Go no es necesario usar esta palabra reservada ya que genera error  
 Este enfoque es consistente con las prácticas recomendadas en Go y evita el error de estilo .	

	this.Id=id
	this.Name=name
	this.CreatedAt=createdAt
	this.Status=status
} */


// El símbolo * se utiliza para denotar un puntero a un tipo de dato.
// En Go, un puntero es una variable que almacena la dirección de memoria de otra variable. Esto proporciona una forma de acceder y 
// modificar indirectamente el valor de la variable original mediante su dirección en memoria.

// Receptor de Valor (User): Cuando se define un método con un receptor de valor (func (u User) Metodo()), Go crea una copia del valor del receptor al
// llamar al método. Los cambios realizados dentro del método no afectan al valor original fuera del método.

// Receptor de Puntero (*User): Cuando se define un método con un receptor de puntero (func (u *User) Metodo()), Go pasa la dirección de memoria del 
// receptor al llamar al método. Esto significa que los cambios realizados dentro del método afectan directamente al valor original del receptor fuera del método.

// func (u *User) AddUser(id int, name string, createdAt time.Time, status bool)  es un método definido para el receptor de puntero *User. Modifica directamente el 
// campo Status del usuario porque u es un puntero a User.

// En resumen, *User en Go representa un tipo puntero que apunta a una estructura User. Utilizar *User como receptor en métodos permite modificar directamente el valor del 
// objeto original, lo cual es útil cuando deseas cambiar el estado de un objeto de manera eficiente sin hacer copias completas de grandes estructuras de datos.


func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) { // con this se refiere a User, no es obligatorio la palabra this se usa por convensión con otros lenguajes, this o self en Go no es necesario usar esta palabra reservada ya que genera error  
	u.Id=id
	u.Name=name
	u.CreatedAt=createdAt
	u.Status=status
}
