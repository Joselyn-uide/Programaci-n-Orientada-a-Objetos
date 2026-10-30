package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

func main() {

	product := Producto{}
	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	fmt.Println(product)
	fmt.Println(product2)
	fmt.Println("El valor del inventario del Producto: ", product2.Nombre, "Tiene un valor del inventario de: ", calcularInventario(product2))
	fmt.Println("El valor del inventario del Producto: ", product.Nombre, "Tiene un valor del inventario de: ", calcularInventario(product))
	fmt.Println("Hay stock disponible? ", product2.Vender(201))

	product3 := NewProducto("mouse", 12.91, 50)
	fmt.Println(product3)
}

func NewProducto(nombre string, precio float64, stock int) Producto {
	return Producto{
		Nombre: nombre,
		precio: precio,
		stock:  0,
	}
}

func calcularInventario(p Producto) float64 {
	return p.precio * float64(p.stock)
}

func (p *Producto) Vender(cantidad int) bool {

	if cantidad < 0 || cantidad >= p.stock {
		return false
	}
	p.stock -= cantidad
	return true
}

func (p *Producto) NuevoStock(cantidad int) {
	if cantidad > 0 {
		p.stock += cantidad
	}
}
