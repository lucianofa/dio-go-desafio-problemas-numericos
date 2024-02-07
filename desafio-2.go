package main

import "fmt"

/* Desafio 2 - exibir todos os números entre 1 e 100 onde
 nos numeros divisíveis por 3 imprimir "Pin"
 nos numeros divisíveis por 5 imprimir "Pan"
 nos numeros divisíveis por 3 e 5 imprimir "Pin Pan"
*/
func main() {

	for i := 1; i <= 100; i++ {

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i, " Pin Pan")
		} else if i % 3 == 0 {
			fmt.Println(i, " Pin")
		} else if i % 5 == 0 {
			fmt.Println(i, " Pan")
		}
		
	}

}
