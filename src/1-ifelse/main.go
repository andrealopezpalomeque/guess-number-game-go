package main

import (
	"fmt"
	"time"
)

func main() {
	//t:= time.Now() //obtiene la hora actual
	//hora := t.Hour() //obtiene la hora de la variable t
	//fmt.Println("La hora es: ", hora)

	//la condicion del if va sin parentesis
	if hora := time.Now().Hour(); hora < 12 { // inicializa la variable hora y la usa solo en el if
		fmt.Println("Buenos dias")

	} else if hora < 17 { 
		fmt.Println("Buenas tardes")

	} else {
		fmt.Println("Buenas noches")
	}
	
}