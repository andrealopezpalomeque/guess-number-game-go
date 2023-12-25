package main

import "fmt"


func main(){

	//var i int

	for i := 1; i <= 10; i++ {
		
		if i == 5 {
			fmt.Println("Se encontro el 5")
			continue // o break
		}
		fmt.Println(i)
	}

	
}

