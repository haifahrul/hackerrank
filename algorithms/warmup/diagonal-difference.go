package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	n := len(arr)

	var primary int32
	var secondary int32

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				primary += arr[i][j]
			}
			if i+j == n-1 {
				secondary += arr[i][j]
			}
		}
	}

	diff := math.Abs(float64(primary - secondary))

	fmt.Println("Primary:", primary)
	fmt.Println("Secondary:", secondary)
	fmt.Println("Diff:", diff)
}
