package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
	fmt.Println(&x)
}

func main() {
	x := 5
	zeroPointer(&x)
	fmt.Println(x)

	zero(x)
	fmt.Println(x)
}
