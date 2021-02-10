package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	var b = []int{3, 2, 1}

	var alice, bob int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}

	var res []int32
	res = append(res, alice)
	res = append(res, bob)

	fmt.Println("Results:", res)

	return
}
