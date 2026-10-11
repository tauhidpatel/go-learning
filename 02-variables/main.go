package main // declares this as an executable program.

import "fmt"

func main() {
	var name string = "Tauhid" // Explicitly declares a string variable
	fmt.Println("Name is:", name)

	var age int = 24 //  Explicitly declares an integer variable
	fmt.Println("Age is:", age)

	var points float64 = 646464.72 //  Explicitly declares a floating-point number
	fmt.Println("Points are:", points)

	isDeveloper := true // declares a variable and lets Go infer its type
	fmt.Println(isDeveloper)

	// GO is statically type, Once age is an int, you can't assign a string to it
	
	// un-comment below line to see the error
	// age = "twenty-four" 
	// this will throw a compiler error
}
