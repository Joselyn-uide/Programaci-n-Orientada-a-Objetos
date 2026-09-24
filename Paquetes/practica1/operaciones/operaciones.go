package operaciones

import "fmt"

func Suma(a, b int) int {
	return a + b
}

func SumarRestar(num1 int, num2 int) (int, int) {
	suma := num1 + num2
	resta := 0
	if num2 > num1 {
		resta = num2 - num1
	} else {
		fmt.Println("El resultado de la resta es 0")
	}
	return suma, resta
}

func Sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
