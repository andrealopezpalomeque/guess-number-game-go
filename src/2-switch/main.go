package main

import (
	"fmt"
	"runtime"
)

func main() {
	// os:= runtime.GOOS //obtiene el sistema operativo

	switch os:= runtime.GOOS; os {
		case "windows":
			fmt.Println("Ejecutando en Windows")
			break //!no es necesario ponerlo, pero si se pone, se sale del switch
		case "linux":
			fmt.Println("Ejecutando en Linux")
		case "darwin":
			fmt.Println("Ejecutando en MacOS")		
		default:
			fmt.Println("No se encontro el sistema operativo")
	}
}