package main

import "fmt"

func main() {
	var estNotas [6][4]float64

	for i := 0; i < 6; i++ {
		fmt.Printf("Estudiante %d:\n", i+1)

		for j := 0; j < 4; j++ {
			fmt.Printf("Nota %d: ", j+1)
			fmt.Scan(&estNotas[i][j])
		}
	}

	var promedio []float64

	sumtotal := 0.0

	for i := 0; i < 6; i++ {

		notas := estNotas[i][:]

		sumanotas := 0.0
		notaAlta := notas[0]
		notaBaja := notas[0]

		for _, notaActual := range notas {
			sumtotal += notaActual
			sumanotas += notaActual

			if notaActual > notaAlta {
				notaAlta = notaActual
			}

			if notaActual < notaBaja {
				notaBaja = notaActual
			}
		}

		prom := sumanotas / 4
		promedio = append(promedio, prom)

		fmt.Printf("\nEstudiante %d\n", i+1)
		fmt.Printf("Promedio: %.2f\n", prom)
		fmt.Printf("Nota más alta: %.2f\n", notaAlta)
		fmt.Printf("Nota más baja: %.2f\n", notaBaja)
	}

	promedioGeneral := sumtotal / 24

	fmt.Printf("\nPromedio general de la clase: %.2f\n", promedioGeneral)
}
