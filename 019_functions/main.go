package main

import "fmt"

// ---------------------------------------------------------
// 1. Function with parameters and one return value
// ---------------------------------------------------------

// add takes two integer parameters:
// a → first number
// b → second number
//
// The function returns one integer value.
// The return type is written after the parameters: int.
func add(a int, b int) int {

	// Add the two numbers and return the result.
	return a + b
}

// ---------------------------------------------------------
// 2. Function with parameters and multiple return values
// ---------------------------------------------------------

// SumAndProduct takes two integers and returns TWO integers.
//
// First return value  → sum
// Second return value → product
//
// (int, int) means that this function returns two values,
// and both values are of type int.
func SumAndProduct(a int, b int) (int, int) {

	// Calculate the sum of a and b.
	sum := a + b

	// Calculate the product of a and b.
	product := a * b

	// Return both values.
	//
	// The first value returned is sum.
	// The second value returned is product.
	return sum, product
}

// ---------------------------------------------------------
// 3. main function
// ---------------------------------------------------------

func main() {

	// Call the add() function with 10 and 20.
	//
	// add(10, 20) returns:
	// 10 + 20 = 30
	//
	// The returned value is stored in the variable "sum".
	sum := add(10, 20)

	// Call SumAndProduct() with 6 and 5.
	//
	// It returns TWO values:
	//
	// Sum     = 6 + 5 = 11
	// Product = 6 * 5 = 30
	//
	// "s" stores the first returned value.
	// "p" stores the second returned value.
	s, p := SumAndProduct(6, 5)

	// Print all three values:
	//
	// sum = 30
	// s   = 11
	// p   = 30
	//
	// Output:
	// 30 11 30
	fmt.Println(sum, s, p)

	// Call SumAndProduct() with 10 and 2.
	//
	// It returns:
	//
	// First value  → 12 (sum)
	// Second value → 20 (product)
	//
	// We only need the sum.
	// Therefore:
	// "Sum" stores the first returned value.
	// "_" receives the second value and ignores it.
	Sum, _ := SumAndProduct(10, 2)

	// Print the sum.
	//
	// Output:
	// 12
	fmt.Println(Sum)
}