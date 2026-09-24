package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("***Bienvenid@s a la clase de Paquetes ****")
	mensaje := saludo.Saludar("Juan")
	fmt.Println(mensaje)

	resultado_suma := operaciones.Suma(20, 10)
	fmt.Println("El resultado de la suma es: ", resultado_suma)

	resultado_sum, resultado_rest := operaciones.SumarRestar(10, 15)
	fmt.Println("El resultado de la suma es: ", resultado_sum)
	fmt.Println("El resultado de la resta es: ", resultado_rest)

	resultado_sumatoria := operaciones.Sumatoria(1, 2, 3, 4, 5)
	fmt.Println("El resultado de la sumatoria es: ", resultado_sumatoria)

}
