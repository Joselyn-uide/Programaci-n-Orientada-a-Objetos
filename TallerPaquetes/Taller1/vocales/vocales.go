package vocales

import "fmt"

func ContarVocales(frase string) {
	count_a := 0
	count_e := 0
	count_i := 0
	count_o := 0
	count_u := 0

	for _, caracter := range frase {
		switch caracter {
		case 'a', 'A', 'á', 'Á', 'ä', 'Ä':
			count_a++
		case 'e', 'E', 'é', 'É', 'ë', 'Ë':
			count_e++
		case 'i', 'I', 'í', 'Í', 'ï', 'Ï':
			count_i++
		case 'o', 'O', 'ó', 'Ó', 'ö', 'Ö':
			count_o++
		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			count_u++
		}
	}

	fmt.Printf("En la frase hay: \n vocales 'a': %d \n vocales 'e': %d \n vocales 'i': %d \n vocales 'o': %d \n vocales 'u': %d\n", count_a, count_e, count_i, count_o, count_u)
}
