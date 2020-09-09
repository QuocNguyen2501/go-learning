package main

import "fmt"

func main(){
	for i := 0; i < 30; i++ {
		bar(i)
	}
	for i := 0; i < 30; i++ {
		foo(i)
	}
}

func bar(i int){
	fmt.Println(fmt.Sprintf(`bar %d`,i))
}

func foo(i int){
	fmt.Println(fmt.Sprintf(`foo %d`,i))
}
