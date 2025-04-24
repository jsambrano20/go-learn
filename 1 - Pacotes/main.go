package main

import (
	"fmt"
	auxiliar "modulo/Auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("dev@gmail.com")
	fmt.Println(erro)
}
