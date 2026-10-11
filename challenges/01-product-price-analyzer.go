package main

import "fmt"

func main() {
	// Ask the user for their name.
	var name string
	
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	
	prices := []int {}
	// Ask them to enter five product prices.
	fmt.Println("Enter 5 product prices")
	for i := range 5 {
		var price int
		fmt.Printf("Price %d: ", i+1)
		fmt.Scan(&price)

		// Append each entered price to the slice.
		prices = append(prices, price)
	}

	total := calculateTotal(prices)
	fmt.Println("Total:", total)

	avg := float64(total) / float64(len(prices))
	fmt.Println("Average price: ", avg)

	if total < 1000 {
		fmt.Println("Budget purchase")
	} else if total >= 100 && total <= 2999 {
		fmt.Println("Moderate purchase")
	} else {
		fmt.Println("Premium purchase")
	}

	var choice string

	fmt.Println("What would you like to display? total / average / both")
	fmt.Scan(&choice)

	switch choice {
	case "total":
	    fmt.Println("Total:", total)
	case "average":
	    fmt.Println("Average:", avg)
	case "both":
	    fmt.Println("Total:", total)
	    fmt.Println("Average:", avg)
	default:
	    fmt.Println("Invalid choice")
	}

}

func calculateTotal(prices []int) int {
	total := 0

	for _, price := range prices {
		total += price
	}

	return total
}