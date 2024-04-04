package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

/*
package main

import "fmt"

func main() {
	for i := 1; i < 20; i += 2 {
		fmt.Println(i)
	}
}
*/
