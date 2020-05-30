package main

import (
	"fmt"
)

func main() {
	slices1 := []string{"Bon", "Bob", "Jerry"}

	// Add a new item into a slices
	slices1 = append(slices1, "May")
	for i, v := range slices1 {
		fmt.Printf("Index: %v ; values: %v\n", i, v)
	}

	fmt.Println("==================================")
	// Add a slice into other
	slices2 := []string{"Peter"}
	slices2 = append(slices2, slices1...)
	for i, v := range slices2 {
		fmt.Printf("Index: %v ; values: %v\n", i, v)
	}

	fmt.Println("==================================")
	// Make a slices with a capacity
	slices3 := make([]int, 0, 6)

	for i := 0; i < 18; i++ {
		slices3 = append(slices3, i)
		fmt.Printf("cap %v; len %v; address: %p\n", cap(slices3), len(slices3), slices3)
	}
	for i, v := range slices3 {
		fmt.Printf("index: %v ; values: %v ; address: %p\n", i, v, slices3)
	}

	fmt.Println("==================================")
	// Remove an element in slice
	slice4 := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Before remove:")
	for i, v := range slice4 {
		fmt.Printf("Index: %v ; values: %v\n", i, v)
	}
	slice4 = append(slice4[:2], slice4[3:]...)
	fmt.Println("After remove:")
	for i, v := range slice4 {
		fmt.Printf("Index: %v ; values: %v\n", i, v)
	}

	fmt.Println("==================================")
	// Remove an element in slice
	slice5 := []int{10, 20, 30, 40, 50, 60}
	slice5 = RemoveElement(slice5, 3)
	for i, v := range slice5 {
		fmt.Printf("Index: %v ; values: %v\n", i, v)
	}
}

func RemoveElement(slices []int, index int) []int {
	slices = append(slices[:index], slices[index+1:]...)
	return slices
}
