package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("\nIngresa una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	var palabras []string
	palabraActual := ""

	for _, char := range frase {
		if char == ' ' {
			if palabraActual != "" {
				palabras = append(palabras, palabraActual)
				palabraActual = ""
			}
		} else {
			palabraActual += string(char)
		}
	}

	if palabraActual != "" {
		palabras = append(palabras, palabraActual)
	}

	conteoPalabras := make(map[string]int)
	for _, word := range palabras {
		conteoPalabras[word]++
	}

	fmt.Println("Resultados del conteo de palabras:")
	for palabra, cantidad := range conteoPalabras {
		fmt.Println(palabra, ":", cantidad)
	}

	max := 0
	masRepetida := ""
	for word, cant := range conteoPalabras {
		if cant > max {
			max = cant
			masRepetida = word
		}
	}
	fmt.Printf("La palabra que más se repite es '%s' con %d repeticiones\n", masRepetida, max)
}
