package main

import (
	"fmt"
)

func main() {
	candles := []int32{3, 2, 1, 3}
	tmp := make(map[int]int)
	var tallest int

	for i := 0; i < len(candles); i++ {
		c := int(candles[i])
		if c > tallest {
			tallest = c
		}
		tmp[c]++
	}

	fmt.Println("Tallest:", tmp[tallest])
}
