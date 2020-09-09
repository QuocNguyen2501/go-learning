package service

import "fmt"

type TrainEngine struct{}

func (t TrainEngine) Assemble() {
	fmt.Println("Assembling components for train engine")
}
