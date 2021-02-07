package main

import "fmt"

func main() {
	test1()
}

func test1() {
	a := []int32{1, 2, 3, 4, 5}
	d := 4 // number of rotation
	// a := []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}
	// d := 10 // number of rotation
	l := len(a)

	tmp := make([]int32, l)

	for i := 0; i < l; i++ {
		if i < d {
			tmp[i] = a[d+i]
		} else {
			tmp[i] = a[i-d]
		}
	}

	fmt.Println("Result  :", tmp)
	fmt.Println("Expected: 77 97 58 1 86 58 26 10 86 51 41 73 89 7 10 1 59 58 84 77")
	fmt.Println("Expected: 5 1 2 3 4")
}
