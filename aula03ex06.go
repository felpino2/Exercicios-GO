package main

import "fmt"

func main() {
	x := -3449395394
	maior := -99999999999999999

	for x != 0 {
		fmt.Println("Qual o número?")
		fmt.Scan(&x)
		if maior < x {
			maior = x
		}
		fmt.Println("O maior é", maior)
	}
}
