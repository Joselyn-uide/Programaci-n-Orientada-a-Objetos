package monedas

import "fmt"

func ConversorMonedas(valor float64, moneda string) float64 {
	var resultado float64
	switch moneda {
	case "euro":
		resultado = valor * 0.8783
	case "libra":
		resultado = valor * 0.7557
	case "won":
		resultado = valor * 1.385
	case "BTC":
		resultado = valor * 0.0000118
	default:
		fmt.Println("Moneda no válida. Por favor, elige entre euro, libra, won o BTC.")
	}
	return resultado
}
