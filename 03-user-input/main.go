package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name: ")
	// reads input from the terminal and stores it in name
	fmt.Scan(&name) 
	// gives Go the memory address of the variable, allowing Scan to write the input into it

	fmt.Println("Enter your age: ")
	fmt.Scan(&age) 

	fmt.Println("Hello", name, "Your age is", age)
}