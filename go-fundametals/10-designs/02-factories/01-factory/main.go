package main

import (
	"go-learning.com/factories/factory"
)

func main() {
	engine := factory.GetEngine("train")
	engine.Assemble()
}
