package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
}

func (usu usuario) salvar() {
	fmt.Println("Salvando usuario")
}

func main() {

	usuario1 := usuario{"João", "Pedro"}
	fmt.Println(usuario1)
	usuario1.salvar()
}
