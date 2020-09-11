package main

import "fmt"

func main(){
	go f("goroutine")

	f("direct")

	go func(msg string){
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
}

func f(from string){
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
