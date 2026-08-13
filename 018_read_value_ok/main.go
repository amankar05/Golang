package main

import "fmt"

func main() {

	// ---------------------------------------------------------
	// 1. Create and initialize a map
	// ---------------------------------------------------------

	// Create a map where:
	// - The key is a string.
	// - The value is an integer.
	//
	// The map contains:
	// "a" → 93
	// "b" → 5
	score := map[string]int{
		"a": 93,
		"b": 5,
	}

	// ---------------------------------------------------------
	// 2. Check whether a key exists in a map
	// ---------------------------------------------------------

	// When accessing a map, Go can return two values:
	//
	// value → the value stored for the given key.
	// ok    → a boolean that tells whether the key exists.
	//
	// Here we are checking for the key "b".
	//
	// Since "b" exists:
	// value = 5
	// ok    = true
	value, ok := score["b"]

	// Print both returned values.
	//
	// Output:
	// 5 true
	fmt.Println(value, ok)

	// ---------------------------------------------------------
	// 3. Using the comma-ok idiom with if
	// ---------------------------------------------------------

	// Check whether the key "c" exists in the map.
	//
	// The if statement creates two variables:
	//
	// value → contains the value associated with "c".
	// ok    → true if "c" exists, false if it does not.
	//
	// The condition "ok" means:
	// "Execute this block only if the key exists."
	if value, ok := score["c"]; ok {

		// This block executes when the key exists.
		fmt.Println(value)

	} else {

		// Since "c" does not exist in the map,
		// ok will be false and this block executes.
		fmt.Println("Not found")
	}

	// ---------------------------------------------------------
	// 4. Create another map
	// ---------------------------------------------------------

	// Create a map containing item names and their prices.
	//
	// Key   → item name
	// Value → price
	prices := map[string]int{
		"xyz": 100,
		"zyx": 200,
	}

	// Initialize total to 0.
	// We will add every item's price to this variable.
	total := 0

	// ---------------------------------------------------------
	// 5. Iterate over a map using range
	// ---------------------------------------------------------

	// range allows us to loop through every key-value pair
	// in the map.
	//
	// item  → receives the key.
	// price → receives the value.
	//
	// First iteration might be:
	// item = "xyz"
	// price = 100
	//
	// Another iteration:
	// item = "zyx"
	// price = 200
	//
	// IMPORTANT:
	// The order of map iteration is NOT guaranteed in Go.
	for item, price := range prices {

		// Print the current key and value.
		fmt.Println(item, price)

		// Add the current price to the total.
		total = total + price
	}

	// Print the total price of all items.
	//
	// 100 + 200 = 300
	fmt.Println(total)
}