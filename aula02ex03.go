package main

import "fmt"

func main() {
	var A float32
	var B float32
	var C float32

	fmt.Println("Qual o valor de A?")
	fmt.Scan(&A)
	fmt.Println("Qual o valor de B?")
	fmt.Scan(&B)
	fmt.Println("Qual o valor de C?")
	fmt.Scan(&C)

	fmt.Println("a área do triângulo retângulo que tem A por base e C por altura.", A*C/2)
	fmt.Println("a área do círculo de raio C.", 3.14*C*C)
	fmt.Println("A Área do trapézio que tem A e B por bases e C por altura.", (A+B)*C/2)
	fmt.Println("a área do quadrado que tem lado B.", B*B)
	fmt.Println("a área do retângulo que tem lados A e B.", A*B)
}
