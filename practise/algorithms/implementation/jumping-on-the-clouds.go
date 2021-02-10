package main

import "fmt"

func main() {

	// ==== SOLUTION 1
	// c := []int{0, 0, 1, 0, 0, 1, 0}
	c := []int{0, 0, 0, 1, 0, 0}
	n := len(c)

	var jumps = 0
	for i := 0; i < n-1; i++ {
		if i+2 < n && c[i+2] == 0 {
			i++
		}
		jumps++
	}
	fmt.Println("jumps:", jumps)

	// ==== SOLUTION 2
	// n := 6
	// clouds := []int{0, 0, 0, 0, 1, 0}
	// clouds := []int{0, 1, 0, 0, 0, 1, 0}
	// n := 7
	// clouds := []int{0, 0, 1, 0, 0, 1, 0}
	// position, jumps := 0, 0
	// for position < n-1 {
	// 	if position != n-2 {
	// 		if clouds[position+2] == 0 {
	// 			position += 2
	// 		} else {
	// 			position++
	// 		}
	// 	} else {
	// 		position++
	// 	}
	// 	jumps++
	// }
	// fmt.Println("jumps:", jumps)
}
