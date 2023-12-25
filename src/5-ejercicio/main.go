//generar un numero aleatorio entre 0 y 100

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//creo una semilla para que no me genere siempre el mismo numero
	//rand.Seed(time.Now().UnixNano())
	
	//fmt.Println(rand.Intn(100))

	jugar()
	jugarNuevamente()
}


func jugar(){
	
	numAleatorio := rand.Intn(100) + 1 //para que no me genere el 0

	var numIngresado int
	var intentos int
	const maxIntentos = 5

	//bucle for donde se intenta adivinar el numero
	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un numero (intentos restantes %d): ", maxIntentos - intentos+1)
		fmt.Scanln(&numIngresado) //& para pasar la direccion de memoria de la variable // scanln lee hasta que se ingrese un enter
		

		if numIngresado == numAleatorio {
			fmt.Println("Ganaste")
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a advinar es MAYOR")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es MENOR")
		}
	}

	fmt.Println("Perdiste. El numero era: ", numAleatorio)


}

func jugarNuevamente(){

	var respuesta string

	fmt.Println("Desea jugar nuevamente? (s/n)")
	fmt.Scanln(&respuesta)

	switch respuesta {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Respuesta invalida")
		jugarNuevamente()	
	}	
}