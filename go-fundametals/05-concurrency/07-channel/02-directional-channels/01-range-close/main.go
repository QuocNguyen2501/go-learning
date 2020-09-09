package main

import "fmt"

func main(){
	c :=make(chan int)

	go foo(c)

	for v:=range c {
		fmt.Println(v)
	}

	fmt.Println("exit...")
}

func foo(c chan<- int){
	for i:=0;i<5;i++{
		c<-i
	}
	close(c)
}
