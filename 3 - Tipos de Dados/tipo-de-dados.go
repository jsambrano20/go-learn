package main

import (
	"errors"
	"fmt"
)

func main() {

	//NUMEROS
	var numero int = 123
	fmt.Println(numero)

	var numero2 uint32 = 2000 //unsygned int
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 1000
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 100
	fmt.Println(numero4)

	//float32 e 64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345.67
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	//FIM NUMEROS

	//BOOLEANO
	var booleano1 bool
	fmt.Println(booleano1)
	//FIM BOOLEANO

	//error
	var error error = errors.New("Erro interno")
	fmt.Println(error)
	//FIM BOOLEANO
}
