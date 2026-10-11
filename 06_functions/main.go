package main

import "fmt"

// As programs grow, we want to split logic into reusable functions, just like methods in C#.
func main() {
	prices := []int{100, 250, 399, 500, 650}

	// calculateTotal(prices) calls the function.
	total := calculateTotal(prices)
	fmt.Println("Total:", total)
}

// declares a function.
// prices []int means the function accepts a slice of integers.
// int after the parentheses declares the return type.
// return total sends the result back to the caller.
func calculateTotal(prices []int) int {
	total := 0

	for _, price := range prices {
		total += price
	}

	return total
}