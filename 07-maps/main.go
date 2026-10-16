package main

import "fmt"

func main() {
	// A map stores key-value pairs, so you can look up a value using a key.
	prices := map[string]int{
		"Keyboard": 1500,
		"Mouse":    1500,
		"Monitor":  1500,
	}

	fmt.Println(prices)
	fmt.Println("Keyboard price:", prices["Keyboard"])

	prices["Mouse"] = 1000
    prices["Headphones"] = 2500
	fmt.Println("Updated prices:", prices)

	delete(prices, "Monitor")
    fmt.Println("After deletion:", prices)

	var product string
	fmt.Println("Enter a product name:")
	fmt.Scan(&product)

	price, exists := prices[product]
	if exists {
		fmt.Printf("The price of %s is %d\n", product, price)
	} else {
		fmt.Println("Product not found")
	}
}