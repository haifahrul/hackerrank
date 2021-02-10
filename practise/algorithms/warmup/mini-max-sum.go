package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{1, 3, 5, 7, 9}
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)

	var values []int

	for i := 0; i < n; i++ {
		var v int
		for j := 0; j < n; j++ {
			if i != j {
				v += arr[j]
			}
		}
		values = append(values, v)
	}
	sort.Ints(values)

	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[n-1])

}
