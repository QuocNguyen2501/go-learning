package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){
	var counter int32 = 0
	const gs = 100

	wg := &sync.WaitGroup{}
	wg.Add(gs)


	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt32(&counter,1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count: ", counter)
}
