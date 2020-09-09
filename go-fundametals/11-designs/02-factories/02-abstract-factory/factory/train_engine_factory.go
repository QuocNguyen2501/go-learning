package factory

import (
	"fmt"

	"go-learning.com/abstract-factory/service"
)

type TrainEngineFactory struct{}

func (t *TrainEngineFactory) NewEngine(engineType string) (service.Engine, error) {
	switch engineType {
	case "COAL_TRAIN_ENGINE":
		return new(service.CoalTrainEngine), nil
	case "ELECTRIC_TRAIN_ENGINE":
		return new(service.ElectricTrainEngine), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Train Engine of type %s undefined", engineType))
	}
}
