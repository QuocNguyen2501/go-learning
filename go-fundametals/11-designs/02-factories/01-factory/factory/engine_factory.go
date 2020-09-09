package factory

import (
	"log"

	"go-learning.com/factories/service"
)

func GetEngine(engineType string) service.Engine {
	switch engineType {
	case "car":
		return &service.CarEngine{}
	case "train":
		return &service.TrainEngine{}
	default:
		log.Printf("type undefined")
		return nil
	}
}
