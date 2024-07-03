package webserver

import "net/http" // paquete para lo que tiene que ver con http

func MiWebServer() {
	http.HandleFunc("/", home) // se coloca un manejador que recibe la petición , que dice que cuando todo lo que sea / ejecute la función home
	// todo lo que tiene que ver con servidores web se maneja con puertos, aca vamos a trabajar con el puerto 3000
	http.ListenAndServe(":3000", nil) // función escuchar y servir en puerto 3000
}

// sintaxis basica de web server cuando una función tiene que mostar contenido, cuando tiene que procesar una petición de http recibe
// estos parametros:
func home(writer http.ResponseWriter, request *http.Request) { // request *http.Request  puntero a http.Request

	http.ServeFile(writer, request, "./webserver/index.html") // ServeFile es una fucnión que permite servir a nuestra web un archivo y ese archivo va a ser nuestro index
	// index html es el archivo que vamos a servir
}
