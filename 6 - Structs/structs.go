package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	u := usuario{"Joao", 20, endereco{"Rua 1", 123}}

	var u2 usuario
	u2.nome = "Lucas"
	u2.idade = 25

	fmt.Println(u)
	fmt.Println(u2)

	var u3 usuario = usuario{nome: "Maria"}
	fmt.Println(u3)

}
