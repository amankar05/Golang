package main

import "fmt"

// divide takes two integer parameters: a and b.
//
// The function has TWO named return values:
//   john   → first return value, of type int
//   sanman → second return value, of type int
//
// Because the return values are named, Go automatically creates
// the variables "john" and "sanman" inside this function.
func divide(a int, b int) (john int, sanman int) {

	// Store the result of a / b in the named return variable "john".
	//
	// For divide(10, 10):
	// john = 10 / 10
	// john = 1
	john = a / b

	// Store the result of a + b in the named return variable "sanman".
	//
	// For divide(10, 10):
	// sanman = 10 + 10
	// sanman = 20
	sanman = a + b

	// A bare "return" returns the current values of all
	// named return variables.
	//
	// So this is equivalent to:
	//
	// return john, sanman
	//
	// At this point:
	// john   = 1
	// sanman = 20
	return
}

func main() {

	// Call the divide() function with:
	// a = 10
	// b = 10
	//
	// The function returns two values:
	//
	// First value  → john   = 1
	// Second value → sanman = 20
	//
	// j receives the first returned value.
	// s receives the second returned value.
	j, s := divide(10, 10)

	// Print both returned values.
	//
	// Output:
	// 1 20
	fmt.Println(j, s)
}