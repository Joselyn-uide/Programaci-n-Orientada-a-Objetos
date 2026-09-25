package main

import (
	"bufio"
	"fmt"
	"os"
	"taller/monedas"
	"taller/vocales"
)

func main() {
	var moneda_usr float64
	var tipomoneda_usr string
	var conversion_moneda float64
	for {
		fmt.Print("Ingrese el valor en dólares: ")
		fmt.Scan(&moneda_usr)
		if moneda_usr < 0 {
			fmt.Println("El valor ingresado no puede ser negativo. Intenta de nuevo.")
		} else {
			break
		}
	}

	for {
		fmt.Print("Ingrese el tipo de moneda (Euros, LB, Won, BTC): ")
		fmt.Scan(&tipomoneda_usr)

		conversion_moneda = monedas.ConversorMonedas(moneda_usr, tipomoneda_usr)
		if conversion_moneda != 0 {
			break
		}
	}
	fmt.Scanln()

	fmt.Printf("El valor de %.2f dólares en %s es: %.2f\n", moneda_usr, tipomoneda_usr, conversion_moneda)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nIngresa una frase: ")
	scanner.Scan()
	frase := scanner.Text()
	vocales.ContarVocales(frase)
}
