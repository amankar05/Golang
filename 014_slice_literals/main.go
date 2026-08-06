package main

import "fmt"

func main() {

	// Create a slice of strings.
	// Slices are dynamic, meaning they can grow or shrink in size.
	result := []string{"Aman", "Kar", "GO", "JAVA"}

	// Print:
	// 1. The complete slice.
	// 2. The first element (index 0).
	// 3. The last element using len(slice)-1.
	fmt.Println(result, result[0], result[len(result)-1])

	// Update the value at index 1.
	// Slices are mutable, so elements can be modified after creation.
	result[1] = "Learn"

	// Print the updated slice.
	// Output: [Aman Learn GO JAVA]
	fmt.Println(result)
}