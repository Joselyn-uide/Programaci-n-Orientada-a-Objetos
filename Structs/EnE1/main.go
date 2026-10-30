package main

import "fmt"

type Persona struct {
	nombre string
}

func (P Persona) GetNombre() string {
	return P.nombre
}

func (P *Persona) SetNombre(nombre string) {
	P.nombre = nombre
}

func main() {
	persona1 := Persona{}

	persona1.SetNombre("Joss")
	fmt.Println("Su nombre es: ", persona1.GetNombre())

	persona1.SetNombre("Ariel")
	fmt.Println("Ahora, su nombre es: ", persona1.GetNombre())

}
