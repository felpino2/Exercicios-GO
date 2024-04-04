package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {

	var number01, number02 int

	fmt.Println("Number one")
	fmt.Scan(&number01)
	fmt.Println("Number two")
	fmt.Scan(&number02)

	fmt.Println(sum(number01, number02))
}
