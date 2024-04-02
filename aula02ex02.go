package main

import "fmt"

func aula02ex2() {

	var number01, number02 float32
	var option int

	fmt.Println("Insira dois operandos")
	fmt.Scan(&number01)
	fmt.Scan(&number02)

	fmt.Println("Qual a opção 1 = +, 2 = -, 3 = /, 4 = *")
	fmt.Scan(&option)
	if option == 1 {
		fmt.Println(number02 + number01)
	} else if option == 2 {
		fmt.Println(number01 - number02)
	} else if option == 3 {
		fmt.Println(number01 / number02)
	} else if option == 4 {
		fmt.Println(number01 * number02)
	} else {
		fmt.Print("Inválido")
	}
}
