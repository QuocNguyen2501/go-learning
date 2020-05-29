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
	switch name {
	case "Job":
		salutation += "Dr "
		fallthrough
	case "Bob":
		salutation += "Mr "
	case "Mary":
		salutation += "Mrs "
	default:
		salutation += "Due "
	}
	return salutation
}
