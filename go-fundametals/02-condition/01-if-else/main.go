package main

import (
	"fmt"
)

func main() {
	i := 2
	if i%2 == 0 {
		fmt.Printf("%v is even number", i)
	} else {
		fmt.Printf("%v is odd number", i)
	}
}
