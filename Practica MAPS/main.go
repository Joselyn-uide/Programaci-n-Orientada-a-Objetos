package main

import "fmt"

func totalvisitas(visitas map[string]int) int {
	total := 0
	for _, valor := range visitas {
		total += valor
	}
	return total
}

func main() {
	visitas := make(map[string]int)

	visitas["Inicio"] = 10
	visitas["Noticias"] = 20
	visitas["Deportes"] = 25
	visitas["Hogar"] = 80

	fmt.Println("Las categorias y sus valores inciales son: ")
	for key, value := range visitas {
		fmt.Println(key, ":", value)
	}

	visitas["Galería"] = 20
	visitas["Inicio"] = 200

	fmt.Println("Las categorias con sus valores cambiados son: ")
	for key, value := range visitas {
		fmt.Println(key, ":", value)
	}

	numvisitas := totalvisitas(visitas)
	fmt.Println("El total de visitas es:", numvisitas)
}
