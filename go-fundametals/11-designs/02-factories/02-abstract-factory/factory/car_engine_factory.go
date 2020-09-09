package factory

import (
	"fmt"

	"go-learning.com/abstract-factory/service"
)

type CarEngineFactory struct{}

func (c *CarEngineFactory) NewEngine(engineType string) (service.Engine, error) {
	switch engineType {
	case "SPORT_CAR_ENGINE":
		return new(service.SportCarEngine), nil
	case "STREET_CAR_ENGINE":
		return new(service.StreetCarEngine), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Car Engine of type %s undefined", engineType))
	}
}
