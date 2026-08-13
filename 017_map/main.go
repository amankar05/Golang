package main

import "fmt"

func main() {

	// ---------------------------------------------------------
	// 1. Creating and initializing a map
	// ---------------------------------------------------------

	// A map stores data in KEY-VALUE pairs.
	//
	// Syntax:
	// map[KeyType]ValueType
	//
	// In this example:
	// - Key type   = string
	// - Value type = int
	//
	// So each name (string) is associated with an age (int).
	//
	// Example:
	// "Aman" → 23
	// "Kar"  → 22
	// "Sahi" → 23
	ages := map[string]int{
		"Aman": 23,
		"Kar":  22,
		"Sahi": 23,
	}

	// Access a value from the map using its key.
	//
	// ages["Aman"] means:
	// "Find the value associated with the key 'Aman'."
	//
	// Output:
	// 23
	//
	// len(ages) returns the total number of key-value pairs
	// currently stored in the map.
	//
	// Output:
	// 3
	fmt.Println(ages["Aman"], len(ages))


	// ---------------------------------------------------------
	// 2. Declaring a map without initializing it
	// ---------------------------------------------------------

	// Declare a map variable without assigning any actual map
	// to it.
	//
	// The map is currently nil.
	//
	// A nil map can be read from, but we cannot add new
	// key-value pairs to it.
	var scores map[string]int

	// Reading from a nil map is allowed in Go.
	//
	// scores prints:
	// map[]
	//
	// scores["a"] tries to find the key "a".
	// Since the map is nil and the key doesn't exist,
	// Go returns the zero value of the value type.
	//
	// The value type here is int.
	// The zero value of int is 0.
	//
	// Therefore:
	// scores["a"] → 0
	fmt.Println(scores, scores["a"])


	// ---------------------------------------------------------
	// 3. Initializing a map using make()
	// ---------------------------------------------------------

	// make() creates and initializes the map.
	//
	// This is important because now the map is no longer nil
	// and we can add key-value pairs to it.
	//
	// Syntax:
	// make(map[KeyType]ValueType)
	scores = make(map[string]int)

	// Add a new key-value pair to the map.
	//
	// "Math" is the key.
	// 90 is the value.
	//
	// The map now contains:
	// Math → 90
	scores["Math"] = 90

	// Print the complete map.
	//
	// Output:
	// map[Math:90]
	fmt.Println(scores)


	// ---------------------------------------------------------
	// 4. Creating another map using make()
	// ---------------------------------------------------------

	// Create an empty map where:
	// - Keys are strings.
	// - Values are integers.
	//
	// make() initializes the map so that we can add values
	// to it immediately.
	//
	// Note:
	// The variable name "scoores" is a spelling mistake.
	// A better name would be "scores2" or "subjects".
	scoores := make(map[string]int)

	// Add a key-value pair to the map.
	//
	// Key   = "English"
	// Value = 99
	scoores["English"] = 99

	// Print the complete map.
	//
	// Output:
	// map[English:99]
	fmt.Println(scoores)


	// ---------------------------------------------------------
	// 5. Creating a map with string keys and string values
	// ---------------------------------------------------------

	// Create and initialize a map where:
	// - Key type   = string
	// - Value type = string
	//
	// This means both the key and the value must be strings.
	//
	// Example:
	// "U1" → "AM"
	// "U2" → "AN"
	// "U3" → "KAR"
	user := map[string]string{
		"U1": "AM",
		"U2": "AN",
		"U3": "KAR",
	}

	// Print the complete map.
	//
	// Output:
	// map[U1:AM U2:AN U3:KAR]
	fmt.Println(user)


	// ---------------------------------------------------------
	// 6. Deleting a value from a map
	// ---------------------------------------------------------

	// delete() removes a key-value pair from a map.
	//
	// Syntax:
	// delete(mapName, key)
	//
	// Here:
	// - user is the map from which we want to delete.
	// - "U3" is the key we want to remove.
	//
	// Before deletion:
	// U1 → AM
	// U2 → AN
	// U3 → KAR
	//
	// After deletion:
	// U1 → AM
	// U2 → AN
	delete(user, "U3")

	// Print the map after deleting the "U3" key.
	//
	// Output:
	// map[U1:AM U2:AN]
	fmt.Println(user)
}