package main

import (
	"fmt"
	"time"
)

func main() {
	var s = "07:05:45PM"

	t, err := time.Parse("15:05:04PM", s)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	fmt.Println("Result:", t.Format("15:05:04"))
}
