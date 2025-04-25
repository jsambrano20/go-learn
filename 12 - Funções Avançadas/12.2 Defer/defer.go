package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}
func main() {
	fmt.Println("Arquivo Defer")
	defer funcao1()
	funcao2()
}
