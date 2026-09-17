package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Ingresa un número entero positivo n: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}

	if numero%2 == 0 {
		fmt.Println("El número es Par")
	} else {
		fmt.Println("El número es Impar")
	}

	switch {
	case numero >= 1 && numero <= 5:
		fmt.Println("Número pequeño")
	case numero >= 6 && numero <= 10:
		fmt.Println("Número mediano")
	case numero > 10:
		fmt.Println("Número grande")
	}
}
