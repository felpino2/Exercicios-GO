package main

import "fmt"

func main() {
	var x float32
	var y float32

	fmt.Print("Qual a coordenada do x?")
	fmt.Scan(&x)
	fmt.Println("Qual a coordenada do y?")
	fmt.Scan(&y)

	if x == 0 && y == 0 {
		fmt.Println("Origem")
	} else if x == 0 {
		fmt.Println("Eixo X")
	} else if y == 0 {
		fmt.Println("Eixo Y")
	} else if x > 0 && y > 0 {
		fmt.Println("Quadrante 1")
	} else if x > 0 && y < 0 {
		fmt.Println("Quadrante 4")
	} else if x < 0 && y > 0 {
		fmt.Println("Quadrante 2")
	} else if x < 0 && y < 0 {
		fmt.Println("Quadrante 3")
	}
}
