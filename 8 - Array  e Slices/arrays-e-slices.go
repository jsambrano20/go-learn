package main

import "fmt"

func main() {
	fmt.Println("Arquivo arrays e slices")

	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	array1[3] = 40
	array1[4] = 50
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5, 10}
	fmt.Println(slice)
}
