package main

import "fmt"

func main() {

	saludo := saludoHola("Juan")
	fmt.Println(saludo)

	suma, resta := calc(10, 5)
	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)
	
}


func saludoHola(nombre string) string {
	return "Hola " + nombre
}

func calc(num1, num2 int) (sum, res int) {
	sum = num1 + num2
	res = num1 - num2
	//return sum, res
	return //tambien se puede retornar sin especificar las variables
}