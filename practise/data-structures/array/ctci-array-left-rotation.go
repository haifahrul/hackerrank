package main

import "fmt"

func main() {
	test1()
}

func test1() {
	a := []int32{1, 2, 3, 4, 5, 6}
	d := int32(4) // number of rotation
	a2 := []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}
	d2 := int32(10) // number of rotation

	fmt.Println("Input   :  1 2 3 4 5 6")
	fmt.Println("Result  :", calculate(a, d))
	fmt.Println("Expected:  5 6 1 2 3 4")
	fmt.Println("")
	fmt.Println("Input   :  41 73 89 7 10 1 59 58 84 77 77 97 58 1 86 58 26 10 86 51")
	fmt.Println("Result  :", calculate(a2, d2))
	fmt.Println("Expected:  77 97 58 1 86 58 26 10 86 51 41 73 89 7 10 1 59 58 84 77")
}

func calculate(arr []int32, d int32) []int32 {
	var i, l int32
	l = int32(len(arr))
	tmp := make([]int32, l)

	for i = 0; i < l; i++ {
		j := (i + d) % l
		tmp[i] = arr[j]
	}

	return tmp
}
