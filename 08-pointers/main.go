package main

import "fmt"

func main() {
	// A pointer stores the memory address of a value, rather than the value itself.
	age := 24

	fmt.Println("Age:", age) // gets the value stored in age
	fmt.Println("Address of age:", &age) // gets the memory address of age.

	ptr := &age // stores that address in a pointer variable.

	fmt.Println("Pointer:", ptr)
    fmt.Println("Value through pointer:", *ptr) // dereferences the pointer to access the value at that address.

	*ptr = 25
    fmt.Println("Updated age:", age)

	score := 50
	fmt.Println("Score address is:", &score)
	pointerScore := &score
	fmt.Println("Value throug pointer is:", *pointerScore)
	*pointerScore = 100
	fmt.Println("Original score variable:", score)
}