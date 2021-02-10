package main

import "fmt"

func main() {
	arr := []int{-4, 3, -9, 0, 4, 1}
	n := len(arr)

	var positive, negative, zero int
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zero++
		} else if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		}
	}

	fmt.Printf("%.6f\n", float64(positive)/float64(n))
	fmt.Printf("%.6f\n", float64(negative)/float64(n))
	fmt.Printf("%.6f\n", float64(zero)/float64(n))
}
