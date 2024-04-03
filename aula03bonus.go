package main

import "fmt"

func main() {

	var matriz [3][3]int
	var input, somaColuna1, somaColuna2, somaLinha2, somaLinha1, diagonalMain, diagonalSec, contador int

	contador = 0

	//prenchimento da matriz

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ { // preenche coluna primeiro
			fmt.Scan(&input)
			matriz[a][b] = input
			somaLinha2 += matriz[2][b]
			somaColuna2 += matriz[a][2]
		}
	}

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			somaLinha1 += matriz[a][b]

			somaColuna1 += matriz[b][a]

			contador++
			if contador == 2 {
				contador = 0
				if somaLinha1 != somaLinha2 {
					fmt.Println("NÃO")
					break
				}
				somaLinha1 = 0
			}

			if a == b {
			}

		}
	}

}
