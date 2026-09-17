package main

import "fmt"

func main() {
	fmt.Println("Bienvenido a la clase de ciclos en Go")
	fmt.Println("*****Bucle Normal*****")

	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	fmt.Println("*******Bucle Infinito*******")
	for {
		fmt.Println("Infinito")
		break
	}

	for rango := range [10]int{} {
		fmt.Println(rango)
	}
}
