package main

import "fmt"

func main() {

	// Create a slice of integers with:
	// Length = 3 → currently contains 3 elements.
	// Capacity = 5 → can hold up to 5 elements before
	// needing to allocate a larger underlying array.
	scores := make([]int, 3, 5)

	// The 3 elements are automatically initialized to 0.
	// Output: [0 0 0]
	fmt.Println(scores)

	// len() returns the number of elements currently in the slice.
	fmt.Println("Length:", len(scores))

	// cap() returns the total storage available in the
	// underlying array before the slice needs to grow.
	fmt.Println("Capacity:", cap(scores))

	// append() adds a new element to the slice.
	// Length increases from 3 to 4.
	// Capacity remains 5 because there is still available space.
	scores = append(scores, 90)

	fmt.Println(scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))

	// Add another element to the slice.
	// Length increases from 4 to 5.
	// Capacity remains 5 because the available capacity is now full.
	scores = append(scores, 80)

	fmt.Println(scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))

	// Create a slice containing programming languages.
	todos := []string{"JAVA", "GO", "Python"}

	// Create another slice containing more programming languages.
	more := []string{"Springboot", "Microservices"}

	// Append all elements from the 'more' slice to the 'todos' slice.
	// The ... spreads the elements of 'more' so that append()
	// receives them as individual arguments.
	todos = append(todos, more...)

	// Output: [JAVA GO Python Springboot Microservices]
	fmt.Println(todos)
}