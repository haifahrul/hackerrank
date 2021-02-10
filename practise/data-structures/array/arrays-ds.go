package main

import "fmt"

func main() {
	arr := []int32{
		1, 4, 3, 2,
	}
	l := int32(len(arr))
	tmp := make([]int32, l)

	for i := int32(0); i < l; i++ {
		tmp[i] = arr[(l-1)-i]
	}

	fmt.Println("Before Reverse:", arr)
	fmt.Println("After Reverse:", tmp)
}
