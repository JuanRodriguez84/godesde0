package files

import (
	"bufio" // bufio sirve para lectura  // no solo sirve para capturar por teclado, sirve para capturar datos de archivos, distintas fuentes de datos
	"fmt"
	"io/ioutil" // libreria   input output util para manejo de archivos para lectura
	"os"

	"github.com/JuanRodriguez84/godesde0/tablaparamanejodearchivos"
)

var fileName string = "./files/txt/tabla.txt"

// graba la tabla en un archivo y crea archivo con la tabla de multiplicar que llega
func GrabaTablaEnArchivo() {
	var texto string = tablaparamanejodearchivos.TablaDeMultiplicarConRetorno()

	archivo, err := os.Create(fileName) // os  sistema operativo metodo create, Create crea un archivo, si el archivo ya existe lo borra y lo crea nuevamente por eso esta SumaTablaEnArchivo
	if err != nil {                     // si hay error
		fmt.Println("Error al crear el archivo" + err.Error())
		return // return sin nada es para que salga de la función
	}

	fmt.Fprintln(archivo, texto) // Fprintln guarda en un archivo el resultado , se dice que se guarde en el buffer creado por create que esta en la variable archivo , texto es lo que almacena en el archivo, es la data
	archivo.Close()              // todo archivo que se use para grabar se debe cerrar, o se use para lectura se debe cerrar
}

// suma tabla toma cualquier archivo que se indique y agrega la tabla de multiplicar al final del archivo
func SumaTablaEnArchivo() {
	var texto string = tablaparamanejodearchivos.TablaDeMultiplicarConRetorno()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644) // OpenFile permite abrir un archivo para hacerlo de tipo append  para concatenrle al final de lo que tenga nuevo contenido
	// os.O_WRONLY|os.O_APPEND -> tags de estritura   o de  concatenación
	// 0644 ->  permisos como en linux como se hace con el chmod
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto) //_, se coloca _ parametro no sirve de nada, WriteString función que permite mediante un string que le pasemos grabar eso en el archivo

	if err != nil {
		fmt.Println("Error durante el writeString" + err.Error())
		return false
	}

	arch.Close()

	return true
}

// Lectura de archivo en GO  con ioutil, lee el archivo completo pero no lo procesa linea por linea
func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo)) // archivo es una SLICE, o conocido como arreglo o coleciones, pero en este caso es SLICE de bytes
}

// metodo de lectura correcta es
func LeoArchivo2daForma() {
	archivo, err := os.Open(fileName) // Open -> para lectura
	if err != nil {
		fmt.Println("Error al leer el archivo en la segunda forma" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo) // como se hace con lectura de teclado, pero en este caso en lugar de recibir datos del teclado se le dice que los datos estan en archivo
	for scanner.Scan() {
		registro := scanner.Text() // devuelve cada linea del archivo
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
