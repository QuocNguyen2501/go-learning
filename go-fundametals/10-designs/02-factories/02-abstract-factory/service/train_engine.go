package service

type TrainEngine interface {
	Engine
	Type() string
}
