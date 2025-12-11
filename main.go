// Author: Bianca Boo
// Version: 1.0.0
// Date: 2025-12-10
// Fileoverview: This program is a simple calculator.

package main

import (
	"fmt"
	"math"
)

func main() {

	// introduction
	fmt.Println("Welcome to my calculator program.")
	fmt.Println("Which operation would you like to perform today?")
	fmt.Println("a. add")
	fmt.Println("b. subtract")
	fmt.Println("c. multiply")
	fmt.Println("d. divide")
	fmt.Println("e. absolute value")
	fmt.Println("f. round")
	fmt.Println("g. raise to an exponent")
	fmt.Println("h. square root")

	// get user choice
	var choice string
	fmt.Print("What operation do you want to choose: ")
	fmt.Scan(&choice)

	// working variables
	var x float64
	var y float64
	var result float64

	// add
	if choice == "a" {
		fmt.Print("Enter first number: ")
		fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		fmt.Scan(&y)

		result = x + y
		fmt.Println(x, "+", y, "=", result)

	// subtract
	} else if choice == "b" {
		fmt.Print("Enter first number: ")
		fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		fmt.Scan(&y)

		result = x - y
		fmt.Println(x, "-", y, "=", result)

	// multiply
	} else if choice == "c" {
		fmt.Print("Enter first number: ")
		fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		fmt.Scan(&y)

		result = x * y
		fmt.Println(x, "*", y, "=", result)

	// divide
	} else if choice == "d" {
		fmt.Print("Enter first number: ")
		fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		fmt.Scan(&y)

		result = x / y
		fmt.Println(x, "/", y, "=", result)

	// absolute value
	} else if choice == "e" {
		fmt.Print("Enter a number: ")
		fmt.Scan(&x)

		result = math.Abs(x)
		fmt.Println("The absolute value of", x, "is", result)

	// round
	} else if choice == "f" {
		fmt.Print("Enter a number: ")
		fmt.Scan(&x)

		result = math.Round(x)
		fmt.Println("The rounded value is", result)

	// exponent
	} else if choice == "g" {
		fmt.Print("Enter base: ")
		fmt.Scan(&x)
		fmt.Print("Enter exponent: ")
		fmt.Scan(&y)

		result = math.Pow(x, y)
		fmt.Println(x, "raised to", y, "is", result)

	// square root
	} else if choice == "h" {
		fmt.Println("In order to calculate the square root, you must supply a non-negative value.")
		fmt.Print("Enter value: ")
		fmt.Scan(&x)

		result = math.Sqrt(x)
		fmt.Println("The square root of", x, "is", result)

	// invalid
	} else {
		fmt.Println("Invalid choice.")
	}

}
