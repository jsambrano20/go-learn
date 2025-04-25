package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo Herança")

	p1 := pessoa{"joao", "pedro", 20, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UEL"}
	fmt.Println(e1)

}
