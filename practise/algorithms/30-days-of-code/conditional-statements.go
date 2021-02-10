package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	if N%2 == 1 {
		fmt.Println("Weird")
	} else if N%2 == 0 && (N > 2 && N <= 5) {
		fmt.Println("Not Weird")
	} else if N%2 == 0 && (N > 6 && N <= 20) {
		fmt.Println("Weird")
	} else if N%2 == 0 && N > 20 {
		fmt.Println("Not Weird")
	}
}
