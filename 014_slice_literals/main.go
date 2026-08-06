package main

import "fmt"

func main() {

	// Create a slice of strings containing two names.
	// Unlike arrays, slices have a flexible size and can grow or shrink.
	result := []string{"aman", "kar"}

	// Print the entire slice.
	// Output: [aman kar]
	fmt.Println(result)
}