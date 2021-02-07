package main

import (
	"fmt"
)

func main() {
	grades := []int{73, 67, 38, 33}
	n := len(grades)

	for i := 0; i < n; i++ {
		if grades[i] < 38 {
			fmt.Println("1", grades[i])
		} else {
			if grades[i]%5 > 2 {
				fmt.Println("2", ((grades[i]/5)+1)*5)
			} else {
				fmt.Println("3", grades[i])
			}
		}
	}
}
