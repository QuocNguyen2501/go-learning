package main

import "fmt"

func main() {
	names := []string{"Bob", "Job", "Mary", "Peter"}
	for _, v := range names {
		salutation := getSalutation(v)
		fmt.Println(salutation + v)
	}
}

func getSalutation(name string) string {
	var salutation string
	switch {
	case name == "Bob", name == "Job", len(name) == 5:
		salutation = "Mr "
	case name == "Mary":
		salutation = "Mrs "
	default:
		salutation = "Due "
	}
	return salutation
}
