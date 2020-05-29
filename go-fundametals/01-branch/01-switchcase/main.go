package main

import "fmt"

func main() {
	names := []string{"Bob", "Mary", "Job", "Peter"}
	for _, v := range names {
		salutation := getSalutation(v)
		fmt.Println(salutation + v)
	}
}

func getSalutation(name string) string {
	switch name {
	case "Bob":
		return "Mr "
	case "Job":
		return "Dr "
	case "Mary":
		return "Mrs "
	default:
		return "Due "
	}
}
