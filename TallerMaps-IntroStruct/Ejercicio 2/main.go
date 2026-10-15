package main

import "fmt"

type Producto struct {
	NombreProducto string
	Precio         float64
	CantidadStock  int
}

func main() {
	producto1 := Producto{
		NombreProducto: "Martillo",
		Precio:         12.50,
		CantidadStock:  15,
	}

	producto2 := Producto{
		NombreProducto: "Taladro",
		Precio:         45.00,
		CantidadStock:  0,
	}

	producto3 := Producto{
		NombreProducto: "Destornillador",
		Precio:         5.75,
		CantidadStock:  20,
	}

	fmt.Println("INFORMACIÓN DE LOS PRODUCTOS")
	listaProductos := []Producto{producto1, producto2, producto3}
	for _, producto := range listaProductos {
		MostrarProducto(producto)
		valorTotal := CalcularValorTotal(producto)
		fmt.Printf("Valor Total: $%.2f\n", valorTotal)
		estadoStock := StockDisponible(producto)
		fmt.Printf("Estado de Stock: %s\n", estadoStock)
		fmt.Println("-----------------------------------------------")
	}
}

func MostrarProducto(product Producto) {
	fmt.Printf("Producto: %s | Precio:$%.2f | Stock: %d unidades \n", product.NombreProducto, product.Precio, product.CantidadStock)
}

func CalcularValorTotal(product Producto) float64 {
	totalProducto := product.Precio * float64(product.CantidadStock)
	return totalProducto
}

func StockDisponible(product Producto) string {
	if product.CantidadStock > 0 {
		return "Disponible"
	} else {
		return "Sin stock"
	}
}
