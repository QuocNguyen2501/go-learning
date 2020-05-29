package main

import (
	"fmt"
)

type Salutation string

func main() {
	var i int
	var k string
	var f bool
	var g Salutation

	getType(i)
	getType(g)
	getType(k)
	getType(f)
}

func getType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
}
