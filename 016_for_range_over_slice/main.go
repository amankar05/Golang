package main

import "fmt"

func main() {

	// Create a slice containing the number of views.
	views := []int{10, 20, 45, 50, 60}

	// Initialize total to 0.
	// We will add each view count to this variable.
	total := 0

	// range iterates over every element in the slice.
	// "_" ignores the index because we only need the value.
	// "v" stores the current value from the slice.
	for _, v := range views {

		// Print the current view count.
		fmt.Println("Views:", v)

		// Add the current value to the total.
		total = total + v
	}

	// Print the sum of all view counts.
	fmt.Println("Total:", total)
}