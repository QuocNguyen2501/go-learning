package main

import "fmt"

func main(){
	f("direct 1 ")

	f("direct 2 ")

	fmt.Scanln()
}

func f(from string){
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
