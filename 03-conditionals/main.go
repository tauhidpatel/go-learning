package main

import "fmt"

func main() {
	var age int
	var role string

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println("You are a child")
	} else if age >= 13 && age <= 17 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}

	fmt.Println("Enter your role:")
	fmt.Scan(&role)

	switch role {
	case "admin":
		fmt.Println("You have full access")
	case "manager":
		fmt.Println("You have access to your team")
	case "developer":
		fmt.Println("You have access to your workspace")
	case "intern":
		fmt.Println("You have limited access")
	default:
		fmt.Println("Unknown acccess")
	}

}