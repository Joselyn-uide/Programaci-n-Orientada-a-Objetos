package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Ingresa un número entero positivo n: ")
	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	switch {
	case numero >= 1 && numero <= 10:
		fmt.Println("Número pequeño")
	case numero >= 11 && numero <= 100:
		fmt.Println("Número mediano")
	case numero > 100:
		fmt.Println("Número grande")
	}
}
