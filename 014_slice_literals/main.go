package main

import "fmt"

func main() {

	// Create a slice of strings.
	// Slices are dynamic, meaning they can grow or shrink in size.
	result := []string{"Aman", "Kar", "GO", "JAVA"}

	// Print:
	// 1. The complete slice.
	// 2. The first element using index 0.
	// 3. The last element using len(result)-1.
	fmt.Println(result, result[0], result[len(result)-1])

	// Update the value at index 1.
	// Slices are mutable, so their elements can be changed.
	result[1] = "Learn"

	// Print the updated slice.
	// Output: [Aman Learn GO JAVA]
	fmt.Println(result)

	// Declare an empty integer slice.
	// The slice currently contains no elements.
	var nums []int

	// Add 10 to the slice using append().
	nums = append(nums, 10)

	// Add multiple values (30 and 49) to the slice.
	nums = append(nums, 30, 49)

	// Print the final slice.
	// Output: [10 30 49]
	fmt.Println(nums)
}