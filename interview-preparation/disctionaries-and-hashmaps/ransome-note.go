package main

import (
	"fmt"
)

func main() {
	checkMagazine(
		[]string{"give", "me", "one", "grand", "today", "night"},
		[]string{"give", "one", "grand", "today"},
	)

	checkMagazine(
		[]string{"two", "times", "three", "is", "not", "four"},
		[]string{"two", "times", "two", "is", "four"},
	)

	checkMagazine(
		[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
		[]string{"ive", "got", "some", "coconuts"},
	)
}

func checkMagazine(magazine []string, note []string) string {
	tmp := map[string]int{}
	replicable := "Yes"

	for _, v := range magazine {
		tmp[v]++
	}

	for _, n := range note {
		if tmp[n] == 0 {
			replicable = "No"
			break
		}
		tmp[n]--
	}

	fmt.Println("Replicable:", replicable)
	fmt.Printf("============================================================\n\n")
	return replicable
}
