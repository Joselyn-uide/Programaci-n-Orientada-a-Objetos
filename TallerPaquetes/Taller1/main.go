package main

import (
	"fmt"
	"taller/monedas"
	"taller/vocales"
)

func main() {
	var moneda_usr float64
	var tipomoneda_usr string
	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&moneda_usr)
	fmt.Print("Ingrese el tipo de moneda (euro, libra, won, BTC): ")
	fmt.Scan(&tipomoneda_usr)

	conversion_moneda := monedas.ConversorMonedas(moneda_usr, tipomoneda_usr)
	if conversion_moneda != 0 {
		fmt.Printf("El valor de %.2f dólares en %s es: %.2f\n", moneda_usr, tipomoneda_usr, conversion_moneda)
	}

	var frase string
	fmt.Print("Ingresa una frase: ")
	fmt.Scan(&frase)
	vocales.ContarVocales(frase)
}
