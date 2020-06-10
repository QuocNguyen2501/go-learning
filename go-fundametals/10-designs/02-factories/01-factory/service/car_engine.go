package service

import "fmt"

type CarEngine struct{}

func (c CarEngine) Assemble() {
	fmt.Println("Assembling components for car engine")
}
