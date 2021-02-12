package main

import "fmt"

func main() {
	s := "aba"
	n := 10
	// s := "a"
	// n := 1000000000000
	// s := "x"
	// n := 970770

	l := len(s)

	times := n / l
	itr := n % l

	count := 0

	for i := 0; i < l; i++ {
		if s[i] == 'a' {
			count++
		}
	}

	fmt.Println("count char a:", count)
	fmt.Println("times:", times)
	fmt.Println("itr:", itr)

	count = count * times
	fmt.Println("count *:", count)

	for i := 0; i < itr; i++ {
		if s[i] == 'a' {
			count++
		}
	}
	fmt.Println("count res:", count)
}
