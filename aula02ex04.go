package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Println("Qual A?")
	fmt.Scan(&A)
	fmt.Println("Qual B?")
	fmt.Scan(&B)
	fmt.Println("Qual C?")
	fmt.Scan(&C)
	fmt.Println("Qual D?")
	fmt.Scan(&D)

	fmt.Println("DIFERENÇA ", A*B-C*D)
}
