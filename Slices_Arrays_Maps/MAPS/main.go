package main

import "fmt"

func main() {

	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	for i := 0; i < 5; i++ {

		fmt.Println("\nElige una actividad:")
		fmt.Println("1. Deportes")
		fmt.Println("2. Videojuegos")
		fmt.Println("3. Cine")
		fmt.Println("4. Música")

		var opcion int
		fmt.Print("Ingresa el número de la actividad: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			votos["deportes"]++

		case 2:
			votos["videojuegos"]++

		case 3:
			votos["cine"]++

		case 4:
			votos["musica"]++

		default:
			fmt.Println("Opción no válida")
		}
	}

	fmt.Println("\nResultados:")

	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad)
	}

	ganador := ganador(votos)

	fmt.Println("\nLa actividad que ganó en votos es:", ganador)
}

func ganador(votos map[string]int) string {

	actividadGanadora := ""
	mayor := 0

	for actividad, cantidad := range votos {

		if cantidad > mayor {
			mayor = cantidad
			actividadGanadora = actividad
		}
	}

	return actividadGanadora
}
