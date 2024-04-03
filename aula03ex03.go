package main

import "fmt"

func main() {
	var i, input int

	fmt.Println("Qual o numero para tabuada?")
	fmt.Scan(&input)

	for i = 1; i <= 10; i++ {
		fmt.Println(i * input)
	}

}
