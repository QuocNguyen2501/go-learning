package bizsrvc

import (
	"log"

	"go-learning.com/abstract-factory/factory"
	"go-learning.com/abstract-factory/service"
)

func BuildFactory(factoryType string) factory.EngineFactory {
	switch factoryType {
	case "CAR_ENGINE_FACTORY":
		return new(factory.CarEngineFactory)
	case "TRAIN_ENGINE_FACTORY":
		return new(factory.TrainEngineFactory)
	default:
		log.Printf("Factory of type %s undefined", factoryType)
		return nil
	}
}

func DeliveringTheEngine(factory factory.EngineFactory, engineType string) service.Engine {
	if factory == nil {
		return nil
	}

	engine, err := factory.NewEngine(engineType)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return engine
}
