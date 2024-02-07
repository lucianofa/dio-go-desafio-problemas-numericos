package main

import "fmt"

// Desafio 1 - exibir todos os números entre 1 e 100 que são divisíveis por 3
func main() {

	for i := 1; i <= 100; i++ {

		if i % 3 == 0 {
			fmt.Println(i, "é divível por 3")
		}
		
	}

}

