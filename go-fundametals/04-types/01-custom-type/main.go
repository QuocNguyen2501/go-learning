package main

import (
	"fmt"
)

type Name string

func main() {
	var name Name
	name = "Will"
	fmt.Printf("\"%T\" has value \"%v\"", name, name)
}
