package main

import "fmt"

func main() {

	var x int
	fmt.Print("Insira <aula><exercicio> ex: 0102")
	fmt.Scan(&x)

	if x == 0201 {
		aula02ex1()
	} else if x == 0202 {
		aula02ex2()
	}
}