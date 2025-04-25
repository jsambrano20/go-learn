package main

import "fmt"

func main() {
	fmt.Println("Arquivo Maps")

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Pedro",
	}
	fmt.Println(usuario)
}
