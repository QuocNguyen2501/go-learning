package main

import (
	"fmt"

	"go-learning.com/abstract-factory/bizsrvc"
)

func main() {
	execute("CAR_ENGINE_FACTORY", "SPORT_CAR_ENGINE")
	execute("CAR_ENGINE_FACTORY", "STREET_CAR_ENGINE")
	execute("TRAIN_ENGINE_FACTORY", "COAL_TRAIN_ENGINE")
	execute("TRAIN_ENGINE_FACTORY", "ELECTRIC_TRAIN_ENGINE")
	execute("TRAIN_ENGINE_FACTORY", "GAS_TRAIN_ENGINE")
	execute("HYBRID_CAR_ENGINE_FACTORY", "SPORT_CAR_ENGINE")
}

func execute(factoryType, engineType string) {
	factory := bizsrvc.BuildFactory(factoryType)
	engine := bizsrvc.DeliveringTheEngine(factory, engineType)
	if engine != nil {
		carEngine, ok := engine.(service.CarEngine)
		if ok {
			fmt.Printf("%s with capacity %d hp and the model is %s\n", carEngine.Assemble(), carEngine.Capacity(), carEngine.Model())
		}

		trainEngine, ok := engine.(service.TrainEngine)
		if ok {
			fmt.Printf("%s with capacity %d hp and the type is %s\n", trainEngine.Assemble(), trainEngine.Capacity(), trainEngine.Type())
		}
	}
}
