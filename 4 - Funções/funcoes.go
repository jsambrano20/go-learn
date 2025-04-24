package main

import "fmt"

func somar(n1 int64, n2 int64) int64 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int64) (int64, int64) {
	soma := n1 + n2
	mult := n1 * n2
	return soma, mult
}

func main() {
	soma := somar(1, 3)
	println(soma)

	calculo, calculo2 := calculosMatematicos(10, 30)
	println(calculo, calculo2)

	var f = func() {
		fmt.Println("Função anônima")
	}
	f()

	var texto = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	texto("Texto")
}
