package main

import "fmt"

func main(){
	c := make(chan int, 1)

	c <- 42
	c <- 32

	fmt.Println(<-c)
	fmt.Println(<-c)
}
