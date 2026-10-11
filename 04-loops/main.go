package main

import "fmt"

func main() {
	// In Go, there's only one looping keyword: for.
	// You can use it for regular loops, while-style loops, and infinite loops.
	for i := 1; i <= 5; i++ {
		// i := 1 — initialize the counter.
		// i <= 5 — continue looping while this condition is true.
		// i++ — increase i by one after each iteration.
		fmt.Println(i)
	}

	var number int = 7
	for i := 1; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number * i)
	}

}
