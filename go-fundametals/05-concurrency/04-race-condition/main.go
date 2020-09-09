package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	var counter int32 = 0
	const gs = 100

	wg := &sync.WaitGroup{}
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(){
			counter ++
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count: ", counter)
}
