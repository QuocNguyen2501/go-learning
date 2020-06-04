package main

import "fmt"

func main() {
	// Create an array with 5 elements
	// but we only set data for 3 first elements
	// by default data will be setted for last elements
	arr1 := [5]string{"Will", "Bob", "Mary"}
	fmt.Printf("%T has values:\n", arr1)
	for i, v := range arr1 {
		fmt.Printf("Index: %v , Value: %v\n", i, v)
	}

	fmt.Println("===================================")

	arr2 := [5]int{10, 20, 30}
	fmt.Printf("%T has values:\n", arr2)
	for i, v := range arr2 {
		fmt.Printf("Index: %v , Value: %v\n", i, v)
	}

	fmt.Println("===================================")

	// Below examples are other way to create a new array
	var arr3 [5]string
	arr3[3] = "Peter"
	arr3[4] = "Will"
	fmt.Printf("%T has values:\n", arr3)
	for i, v := range arr3 {
		fmt.Printf("Index: %v , Value: %v\n", i, v)
	}

	fmt.Println("===================================")

	// We could assign an array to another one
	fmt.Println("Assign an array to another one")
	arr4 := [5]string{"Will", "Bob", "Mary", "Join", "Smith"}
	var arr5 [5]string
	arr5 = arr4 // they have different addresses
	for i, v := range arr5 {
		fmt.Printf("Index: %v , Value: %v\n", i, v)
	}

	// We only remove an item in slice
	// but cannot remove an item in array because we cannot change size of array
}
