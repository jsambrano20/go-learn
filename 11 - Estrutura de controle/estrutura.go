package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")
	numero := 10

	if numero > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor que 10")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que outro numero")
	} else {
		fmt.Println("Menor que outro numero")
	}
}
