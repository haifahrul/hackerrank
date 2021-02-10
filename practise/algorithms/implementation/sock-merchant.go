package main

import (
	"fmt"
)

func main() {

	var n int
	var ar []int

	// n = 9
	// ar = []int{1, 2, 3, 1, 2, 10, 20, 20, 10, 10, 30, 50, 10, 20}
	n = 10
	ar = []int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}

	fmt.Println("- n:", n)
	fmt.Println("- ar:", ar)

	var pairs int
	var tmp = make(map[int]int, 0)

	for i := 0; i < n; i++ {
		tmp[ar[i]]++
	}
	for _, v := range tmp {
		pairs += v / 2
	}
	fmt.Println("tmp", tmp)
	fmt.Println("pairs socks", pairs)

	fmt.Println("======================= SOLVE 2 =======================")
	socks := make([]int, 101)

	for i := 0; i < n; i++ {
		var t int
		t = ar[i]
		socks[t]++
	}
	count := 0

	fmt.Println("socks", socks)

	for _, v := range socks {
		count += v / 2
	}

	fmt.Println("result", count)
}
