package main

import (
	"fmt"
)

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (product *Producto) ActualizarPrecio(nuevoPrecio float64) {
	product.Precio = nuevoPrecio
}

func (product *Producto) ActualizarCantidad(nuevaCantidad int) {
	product.Cantidad = nuevaCantidad
}

func (product Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Producto: %s | Precio: $%.2f | Stock: %d unidades\n",
		product.ID, product.Nombre, product.Precio, product.Cantidad)
}

type Inventario struct {
	listaProductos map[int]Producto
}

func (listP *Inventario) AgregarProducto(product Producto) {
	if listP.listaProductos == nil {
		listP.listaProductos = make(map[int]Producto)
	}
	listP.listaProductos[product.ID] = product
}

func (listP *Inventario) EliminarProducto(id int) {
	delete(listP.listaProductos, id)
	fmt.Println("Producto eliminado")
}

func (listP *Inventario) BuscarProducto(id int) (Producto, bool) {
	producto, existe := listP.listaProductos[id]
	return producto, existe
}

func (listP *Inventario) ListarProductos() {
	if len(listP.listaProductos) == 0 {
		fmt.Println("El inventario de la papelería está vacío.")
	} else {
		for _, producto := range listP.listaProductos {
			producto.MostrarInfo()
		}
	}
}

func main() {
	var tienda Inventario
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
			fmt.Printf("Ingresa el id del producto: ")
			fmt.Scan(&id)
			fmt.Printf("Ingresa el nombre del producto: ")
			fmt.Scan(&nombre)
			fmt.Printf("Ingresa el precio del producto: ")
			fmt.Scan(&precio)
			fmt.Printf("Ingresa la cantidad del producto: ")
			fmt.Scan(&cantidad)
			tienda.AgregarProducto(NuevoProducto(id, nombre, precio, cantidad))
			fmt.Println("Producto agregado")

		case "2":
			var id int
			var nuevoPrecio float64
			fmt.Print("Ingresa el ID del producto: ")
			fmt.Scan(&id)

			product := tienda.listaProductos[id]
			if product.ID != 0 {
				fmt.Print("Ingresa el nuevo precio: ")
				fmt.Scan(&nuevoPrecio)
				product.ActualizarPrecio(nuevoPrecio)
				tienda.AgregarProducto(product)
				fmt.Println("Precio actualizado")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case "3":
			var id int
			var nuevaCantidad int
			fmt.Print("Ingresa el ID del producto: ")
			fmt.Scan(&id)

			product := tienda.listaProductos[id]
			if product.ID != 0 {
				fmt.Print("Ingresa la nueva cantidad: ")
				fmt.Scan(&nuevaCantidad)
				product.ActualizarCantidad(nuevaCantidad)
				tienda.AgregarProducto(product)
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
			product := tienda.listaProductos[id]
			if product.ID != 0 {
				product.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case "6":
			fmt.Println("\n--- LISTA DE PRODUCTOS ---")
			tienda.ListarProductos()

		case "0":
			fmt.Println("Has salido del sistema")
			return

		default:
			fmt.Println("Opción inválida, intenta de nuevo.")
		}
	}
}
