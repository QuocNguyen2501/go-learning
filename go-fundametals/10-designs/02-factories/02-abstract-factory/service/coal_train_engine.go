package service

type CoalTrainEngine struct{}

func (c CoalTrainEngine) Assemble() string {
	return "Assembling components of Coal Train Engine"
}

func (c CoalTrainEngine) Capacity() int {
	return 900000
}
