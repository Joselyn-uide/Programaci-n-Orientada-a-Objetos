package main

import "fmt"

func main() {
	var edad int = 15
	var temperatura float32 = 21.3
	var activo bool = true
	var mensaje string = "Bienvenid@"
	var dato byte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}

	fmt.Printf("Edad: %v---TipoDato: %T\n", edad, edad)
	fmt.Printf("Temperatura: %v---TipoDato: %T\n", temperatura, temperatura)
	fmt.Printf("Activo: %v---TipoDato: %T\n", activo, activo)
	fmt.Printf("Mensaje: %v---TipoDato: %T\n", mensaje, mensaje)
	fmt.Printf("Dato: %v---TipoDato: %T\n", dato, dato)
	fmt.Printf("Días: %v---TipoDato: %T\n", dias, dias)
	fmt.Printf("Números: %v---TipoDato: %T\n", numeros, numeros)

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("Tu edad es: ", edad, "Tu estado es: ", activo)
	fmt.Println("La temperatura es de: ", temperatura)
	fmt.Println("Activo: ", activo)
	fmt.Println("Mensaje: ", mensaje)
	fmt.Println("Datos: ", dato)
	fmt.Println("Días: ", dias)
	fmt.Println("Números: ", numeros)

	var age int
	var decimal float64

	fmt.Println("Ingersa dos valores:")
	fmt.Scanf("%d %f", &age, &decimal)

	fmt.Println("Resultado: ", (float64(age) * decimal * decimal))

}
