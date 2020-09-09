package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("CPU\t\t",runtime.NumCPU())
	fmt.Println("Goroutines\t\t",runtime.NumGoroutine())

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go bar(i,wg)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(i,wg)
	}
	fmt.Println("Goroutines\t\t",runtime.NumGoroutine())
	wg.Wait()
}

func bar(i int,wg *sync.WaitGroup){
	fmt.Println(fmt.Sprintf(`bar %d`,i))
	wg.Done()
}

func foo(i int,wg *sync.WaitGroup){
	fmt.Println(fmt.Sprintf(`foo %d`,i))
	wg.Done()
}
