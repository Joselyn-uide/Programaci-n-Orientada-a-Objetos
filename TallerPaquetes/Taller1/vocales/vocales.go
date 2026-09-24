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
		case 'a', 'A':
			count_a++
		case 'e', 'E':
			count_e++
		case 'i', 'I':
			count_i++
		case 'o', 'O':
			count_o++
		case 'u', 'U':
			count_u++
		}
	}
	fmt.Printf("En la frase hay: \n vocales 'a': %.v \n vocales 'e': %.v \n vocales 'i': %.v \n vocales 'o': %.v \n vocales 'u': %.v\n", count_a, count_e, count_i, count_o, count_u)
}
