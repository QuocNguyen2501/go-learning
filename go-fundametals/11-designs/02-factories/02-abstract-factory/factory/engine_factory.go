package factory

import (
	"go-learning.com/abstract-factory/service"
)

type EngineFactory interface {
	NewEngine(engineType string) (service.Engine, error)
}
