package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// Sample input
	// n = 2
	// loop index 0 = Hacker
	// loop index 1 = Rank

	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		var even, odd string
		for j := 0; j < len(s); j++ {
			if s[j] == 10 {
				continue
			}
			v := string(s[j])

			if j%2 == 0 {
				even = fmt.Sprintf("%s%s", even, string(v))
			} else {
				odd = fmt.Sprintf("%s%s", odd, string(v))
			}
		}
		fmt.Println(even, odd)
	}
}
