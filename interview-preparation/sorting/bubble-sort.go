package main

import "fmt"

func main() {
	var i, j, swap int32
	a := []int32{6, 4, 1}
	n := int32(len(a))

	for i = int32(0); i < n; i++ {
		for j = int32(0); j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				l := a[j]
				n := a[j+1]
				a[j] = n
				a[j+1] = l
				swap++
			}
		}
	}

	fmt.Println("Array is sorted in", swap, "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[n-1])
}
