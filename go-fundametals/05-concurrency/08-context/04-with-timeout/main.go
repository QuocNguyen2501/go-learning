package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main(){
	deadline := time.Duration(2*time.Second)
	ctx, _ := context.WithTimeout(context.Background(),deadline)

	fmt.Println("error check 1: ", ctx.Err())
	fmt.Println("num goroutines 1: ",runtime.NumGoroutine())

	go func(){
		n := 0
		for {
			select {
			case <- ctx.Done():
				fmt.Println("done")
				return
			default:
				n++
				time.Sleep(time.Millisecond * 300)
				fmt.Println("working ",n)
			}
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("error check 2: ", ctx.Err())
	fmt.Println("num goroutines 2: ",runtime.NumGoroutine())


	time.Sleep(time.Second * 1)
	fmt.Println("error check 3: ", ctx.Err())
	fmt.Println("num goroutines 3: ",runtime.NumGoroutine())
}


