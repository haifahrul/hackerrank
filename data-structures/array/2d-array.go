package main

import "fmt"

func main() {
	// arr := [][]int32{
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 2, 4, 4, 0},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }

	// arr := [][]int32{
	// 	{-9, -9, -9, 1, 1, 1},
	// 	{0, -9, 0, 4, 3, 2},
	// 	{-9, -9, -9, 1, 2, 3},
	// 	{0, 0, 8, 6, 6, 0},
	// 	{0, 0, 0, -2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }

	arr := [][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7},
	}

	max := -63

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := arr[i+1][j+1]
			for k := 0; k < 3; k++ {
				sum += arr[i][j+k]
				sum += arr[i+2][j+k]
			}
			if int(sum) > max {
				max = int(sum)
			}
		}
	}

	fmt.Println("Max:", max)
}