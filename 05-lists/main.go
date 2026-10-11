package main

import "fmt"

func main() {
	// Go has slices, which are flexible, resizable collections. You'll use them frequently when building APIs.

	prices := []int{100, 250, 399, 411, 500} //creates a slice of integers.
	total := 0

	fmt.Println("All prices:", prices)
	fmt.Println("First price:", prices[0]) // accesses the first element
	fmt.Println("Number of prices:", len(prices)) // returns the number of elements.

	prices = append(prices, 650) // adds an element and returns the updated slice, which you assign back to prices
    fmt.Println("Updated prices:", prices)

	// range prices, iterates over every element in the slice.
	// index, holds the element's position
	for index, price := range prices { 
		fmt.Println("Index:", index, "| Price:", price)
	}

	// The underscore tells Go to discard a value you don't need. 
	// Go doesn't allow declared local variables to remain unused, so this is useful.
	for _, price := range prices {
    	total += price
	}

	fmt.Println("Total price", total)

	avg := float64(total) / float64(len(prices))
	fmt.Println("Average price: ", avg)

}