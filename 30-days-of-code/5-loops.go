package main

import "fmt"

func main() {
	// sample input
	var n, i, len int32

	n = 2
	len = 10

	for i = 1; i <= len; i++ {
		fmt.Printf("%v x %v = %v\n", n, i, n*i)
	}
}
