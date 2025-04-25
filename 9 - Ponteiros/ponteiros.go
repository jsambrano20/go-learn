package main

import "fmt"

func main() {
	fmt.Println("Arquivo Ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//Ponteiros referencia de memoria
	var var3 int
	var var4 *int

	var3 = 100
	var4 = &var3             // Endereço de memoria
	fmt.Println(var3, *var4) // * desreferenciacao o ponteiro
	fmt.Println(var3, var4)  // Endereço de memoria
}
