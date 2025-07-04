package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" //explicita
	variavel2 := "Variável 2"           //implicita

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
