package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input a single integer:")
	fmt.Scan(&n)

	stair := ""
	for i := 0; i < n; i++ {
		stair = fmt.Sprintf("#%s", stair)
		fmt.Printf("%*s\n", n, stair)
	}
}
