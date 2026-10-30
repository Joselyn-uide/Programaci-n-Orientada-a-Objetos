package main

import (
	"fmt"
)

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) GetID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) GetPrecio() float64 {
	return p.precio
}

func (p Producto) GetCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Producto: %s | Precio: $%.2f | Stock: %d unidades\n",
		p.id, p.nombre, p.precio, p.cantidad)
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.productos[p.GetID()] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.productos, id)
	fmt.Println("Producto eliminado del inventario.")
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
}

func main() {
	tienda := NuevoInventario()
	tienda.AgregarProducto(NuevoProducto(1, "Cuaderno", 2.50, 40))
	tienda.AgregarProducto(NuevoProducto(2, "Lapiz", 0.50, 100))
	tienda.AgregarProducto(NuevoProducto(3, "Borrador", 0.35, 60))
	tienda.AgregarProducto(NuevoProducto(4, "Esfero", 0.60, 80))
	tienda.AgregarProducto(NuevoProducto(5, "Regla", 1.80, 25))

	var opcion string

	for {
		fmt.Println("\n--- MENÚ PAPELERÍA ---")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Actualizar precio")
		fmt.Println("3. Actualizar cantidad")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Buscar producto")
		fmt.Println("6. Listar productos")
		fmt.Println("0. Salir")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			var id, cantidad int
			var nombre string
			var precio float64
			fmt.Print("Ingresa el ID del producto: ")
			fmt.Scan(&id)
			fmt.Print("Ingresa el nombre del producto: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingresa el precio del producto: ")
			fmt.Scan(&precio)
			fmt.Print("Ingresa la cantidad del producto: ")
			fmt.Scan(&cantidad)

			tienda.AgregarProducto(NuevoProducto(id, nombre, precio, cantidad))
			fmt.Println("Producto agregado")

		case "2":
			var id int
			var nuevoPrecio float64
			fmt.Print("Ingresa el ID del producto: ")
			fmt.Scan(&id)

			p := tienda.productos[id]
			if p.GetID() != 0 {
				fmt.Print("Ingresa el nuevo precio: ")
				fmt.Scan(&nuevoPrecio)
				p.SetPrecio(nuevoPrecio)
				tienda.AgregarProducto(p)
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case "3":
			var id int
			var nuevaCantidad int
			fmt.Print("Ingresa el ID del producto: ")
			fmt.Scan(&id)

			p := tienda.productos[id]
			if p.GetID() != 0 {
				fmt.Print("Ingresa la nueva cantidad: ")
				fmt.Scan(&nuevaCantidad)
				p.SetCantidad(nuevaCantidad)
				tienda.AgregarProducto(p)
				fmt.Println("Cantidad actualizada")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case "4":
			var id int
			fmt.Print("Ingresa el ID del producto a eliminar: ")
			fmt.Scan(&id)
			tienda.EliminarProducto(id)

		case "5":
			var id int
			fmt.Print("Ingresa el ID del producto a buscar: ")
			fmt.Scan(&id)

			p := tienda.productos[id]
			if p.GetID() != 0 {
				p.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case "6":
			fmt.Println("\n--- LISTA DE PRODUCTOS ---")
			tienda.ListarProductos()

		case "0":
			fmt.Println("Has salido del sistema.")
			return

		default:
			fmt.Println("Opción inválida, intenta de nuevo.")
		}
	}
}
