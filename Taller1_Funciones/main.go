package main

import "fmt"

//OPCIÓN 1: Promedio de estudiantes
func opcion1() {
	var cantidadEst int
	for {
		fmt.Print("Ingresa la cantidad de estudiantes (en números): ")
		fmt.Scan(&cantidadEst)

		if cantidadEst <= 0 {
			fmt.Println("La cantidad de estudiantes debe ser mayor a 0. Intenta de nuevo.")
		} else {
			break
		}
	}

	var sumaNotas float64 = 0

	for i := 1; i <= cantidadEst; i++ {
		var nota float64
		for {
			fmt.Printf("Ingresa la nota del estudiante %v (de 0 a 100): ", i)
			fmt.Scan(&nota)
			if nota < 0 || nota > 100 {
				fmt.Println("La nota debe estar entre 0 y 100. Intenta de nuevo.")
			} else {
				break
			}
		}
		sumaNotas += nota
	}

	promedio := averageGrade(sumaNotas, cantidadEst)
	fmt.Printf("El promedio es: %.2f\n", promedio)
	if promedio >= 70 {
		fmt.Println("Promedio del curso: Aprobado")
	} else {
		fmt.Println("Reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Excellent performance")
	case promedio >= 80 && promedio <= 89:
		fmt.Println("Good performance")
	case promedio >= 70 && promedio <= 79:
		fmt.Println("Satisfactory performance")
	case promedio < 70:
		fmt.Println("Needs improvement")
	}

}

func averageGrade(totalNotas float64, totalEst int) float64 {
	average := totalNotas / float64(totalEst)
	return average
}

//OPCIÓN 2: Sumar números del 1 al n
func opcion2() {
	var n int
	for {
		fmt.Print("Ingresa un número entero positivo: ")
		fmt.Scan(&n)

		if n <= 0 {
			fmt.Println("El número debe ser mayor a 0. Intenta de nuevo")
		} else {
			break
		}
	}

	var sumaNumeros int
	for i := 1; i <= n; i++ {
		sumaNumeros += i
	}
	fmt.Printf("La suma de los números del 1 al %v es: %v\n", n, sumaNumeros)
}

//OPCIÓN 3: Conversión de Celsius a Fahrenheit
func opcion3() {
	var celsius float64
	fmt.Print("Ingresa la temperatura en grados Celsius (solo el número): ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f grados Celsius son %.2f grados Fahrenheit\n", celsius, fahrenheit)
}

//OPCIÓN 4: Conversión de Fahrenheit a Celsius
func opcion4() {
	var fahrenheit float64
	fmt.Print("Ingresa la temperatura en grados Fahrenheit (solo el número): ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f grados Fahrenheit son %.2f grados Celsius\n", fahrenheit, celsius)
}

func main() {
	for {
		fmt.Println("Menú")
		fmt.Println("1. Promedio de estudiantes")
		fmt.Println("2. Sumar de números del 1 al n")
		fmt.Println("3. Conversión de Celsius a Fahrenheit")
		fmt.Println("4. Conversión de Fahrenheit a Celsius")
		fmt.Println("Ingresa 0 o 'salir' para cerrar el menú")
		fmt.Print("Escribe una opción: ")

		var opcion string
		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Chaooo")
			break
		}
		switch opcion {
		case "1":
			opcion1()
		case "2":
			opcion2()
		case "3":
			opcion3()
		case "4":
			opcion4()
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}
