package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println("A area é", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 3.14 * c.raio * c.raio
}

func main() {
	r := retangulo{altura: 10, largura: 5}
	c := circulo{raio: 10}
	escreverArea(r)
	escreverArea(c)

}
