package monedas

import "fmt"

func ConversorMonedas(valor float64, moneda string) float64 {
	var resultado float64
	switch moneda {
	case "Euros", "euro":
		resultado = valor * 0.8783
	case "LB", "lb", "libra":
		resultado = valor * 0.7557
	case "Won", "won":
		resultado = valor * 1.385
	case "BTC", "btc":
		resultado = valor * 0.0000118
	default:
		fmt.Println("Moneda no válida. Por favor, elige entre Euros, LB, Won o BTC.")
	}
	return resultado
}
