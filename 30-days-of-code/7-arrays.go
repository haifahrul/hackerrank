package main

import "fmt"

func main() {
	// Sample input
	// 1, 4, 3, 2

	arr := []int32{1, 4, 3, 2}
	n := len(arr)

	for i := 0; i < n; i++ {
		fmt.Printf("%v ", arr[n-1-i])
	}
}
