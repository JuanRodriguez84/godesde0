package defer_panic_recover

import (
	//	"os"
	"fmt"
	"log" // permite hacer lo mismo que un println, con la diferencia que el log antepone la fecha y la hora de ejecución
)

func VemosDefer() {
	// Defer es una instrucción que permite configurar cual va a ser la ultima instrucción en ejecutarse cuando salga de la función
	// se usa cuando hay una salida abupta por un error
	// Defer permite configurar
	// ejemplo
	// defer archivo.close()    -> Go asegura que cuando termina la función  va a cerrar el archivo, ya sea por un error o sea lo que pase

	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este es el mensaje final") // es el ultimo mensaje que se mostraría, el defer hace que se ejecute de ultima
	// se pueden hacer varios defer en una función
	fmt.Println("Este es el segundo final")

}

func EjemploPanic() {
	// panic aborta el programa con un mensaje que envia a consola y deja información de lo que abortó

	a := 1

	if a == 1 {
		panic("Se encontró el valor 1")
	}

}

func EjemploRecover() {
	// recover permite recuperarme de un panic, en lugar de dejar que aborte me recupero y luego manejo el error de esa manera , se pueden hacer operaciones en base a
	// los panic
	// antes de abortar se pueden cerrar archivos yu aplicar logica con este recover y luego si abortar

	defer func() { // El defer solo acepta una instrucción asi que si se quiere colocar muchas instrucciones se usa una función anonima
		reco := recover() // el recover se usa con defer si o si , sin el recover no tiene forma de capturar el panic, por que el defer obliga al compilador a ejecutar si o si el codigo aunque exista un panic que abortó el sistema
		// cuando el recover detecta un panic en la variable reco hay un valor que no es nil, por que si no hubo panico el defer va a ejecutarlo igual
		if reco != nil { // si es distinto de nil significa wue hubo un panic
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco) // fatal ->  equivalente a printf y ademas aborta el sistema  , %v  -> objeto de tipo variante, osea que va arecibir cualquier parametro  de cualquier tipo
		}
	}() // las funciones anonimas terminan de esta manera ->    }()
	a := 1

	if a == 1 {
		panic("Se encontró el valor 1")
	}

}
