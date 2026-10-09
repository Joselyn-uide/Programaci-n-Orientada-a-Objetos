package main

import "fmt"

func main() {

	actividadesVotos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("VOTACIÓN")
	for i := 0; i < 5; i++ {
		var opcionUsr int
		for {
			fmt.Println("1. Deportes")
			fmt.Println("2. Videojuegos")
			fmt.Println("3. Cine")
			fmt.Println("4. Música")

			fmt.Printf("Voto %d - Ingresa el número de la actividad de su preferencia: ", i+1)
			fmt.Scan(&opcionUsr)

			if opcionUsr >= 1 && opcionUsr <= 4 {
				break
			} else {
				fmt.Println("Opción no válida. Por favor, elige un número del 1 al 4.")
			}
		}

		switch opcionUsr {
		case 1:
			actividadesVotos["deportes"]++
		case 2:
			actividadesVotos["videojuegos"]++
		case 3:
			actividadesVotos["cine"]++
		case 4:
			actividadesVotos["musica"]++
		default:
			fmt.Println("Opción no válida")
		}
	}
	fmt.Println("\nResultados:")

	for actividad, cantidad := range actividadesVotos {
		fmt.Println(actividad, ":", cantidad)
	}

	ganador := ganador(actividadesVotos)
	fmt.Println("\nLa actividad que ganó en votos es:", ganador)
}

func ganador(actividadesVotos map[string]int) string {

	actividadGanadora := ""
	mayor := 0

	for actividad, cantidad := range actividadesVotos {

		if cantidad > mayor {
			mayor = cantidad
			actividadGanadora = actividad
		}
	}

	return actividadGanadora
}
