package main

import "fmt"

func main() {
	names := []string{"Bob", "Mary", "Peter", "Job"}
	for i := 0; i < len(names); i++ {
		salutation := getSalutation(names[i])
		fmt.Println(salutation + names[i])
	}
}

func getSalutation(name string) string {
	switch name {
	case "Bob", "Job":
		return "Mr "
	case "Mary":
		return "Mrs "
	default:
		return "Due "
	}
}
