package main

import "fmt"

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	var sum int64
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	fmt.Println("SUM:", sum)
	return
}
