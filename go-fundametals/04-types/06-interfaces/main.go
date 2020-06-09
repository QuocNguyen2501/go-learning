package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type CarEngine struct{}

func (e CarEngine) Start() {
	fmt.Println("Car engine is started")
}

func (e CarEngine) Stop() {
	fmt.Println("Car engine is stopped")
}

type TrainEngine struct{}

func (e TrainEngine) Start() {
	fmt.Println("Train engine is started")
}

func (e TrainEngine) Stop() {
	fmt.Println("Train engine is stopped")
}

func Starting(e Engine) {
	e.Start()
}

func Stopping(e Engine) {
	e.Stop()
}

func main() {
	trainEngine := TrainEngine{}
	carEngine := CarEngine{}
	engines := []Engine{trainEngine, carEngine}
	for _, e := range engines {
		Starting(e)
		Stopping(e)
	}
}
