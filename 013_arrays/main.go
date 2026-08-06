package main

import "fmt"

func main() {

	// Declare an integer array that can store exactly 3 elements.
	// By default, all elements are initialized to 0.
	var arr [3]int

	// Assign values to each index of the array.
	arr[0] = 10
	arr[1] = 50
	arr[2] = 40

	// Print the complete array.
	// Output: [10 50 40]
	fmt.Println(arr)

	// Declare and initialize another array with 5 integer values.
	// The size (5) is fixed and must match the number of elements.
	res := [5]int{2, 3, 4, 6, 8}

	// len() returns the total number of elements in the array.
	// Output: 5
	fmt.Println(len(res))
}
