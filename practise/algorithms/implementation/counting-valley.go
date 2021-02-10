package main

import "fmt"

func main() {
	steps := 16
	path := "UDDDUDUUUDDDUDUU"

	depth := 0
	valley := 0
	for i := 0; i < steps; i++ {
		if path[i] == 85 { //U
			depth++
		}
		if path[i] == 68 { //D
			depth--
		}

		if depth == 0 && path[i] == 85 {
			valley++
		}
	}
	fmt.Println(valley)
}
