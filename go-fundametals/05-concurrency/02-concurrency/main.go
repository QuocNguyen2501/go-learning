package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("Goroutines\t\t",runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go bar(i)
	}
	for i := 0; i < 10; i++ {
		go foo(i)
	}
	fmt.Println("Goroutines\t\t",runtime.NumGoroutine())
}

func bar(i int){
	fmt.Println(fmt.Sprintf(`bar %d`,i))
}

func foo(i int){
	fmt.Println(fmt.Sprintf(`foo %d`,i))
}
