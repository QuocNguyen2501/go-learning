package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		if i < 100 {
			fmt.Println(i)
			i++
			continue
		}
		break
	}
}
