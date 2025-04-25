package main

import "fmt"

// 1 1 2 3 5 8 13 21 34 55 89 144
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
func main() {
	fmt.Println("Funções recursivas")

	fmt.Println(fibonacci(10))

	for i := int(0); i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
