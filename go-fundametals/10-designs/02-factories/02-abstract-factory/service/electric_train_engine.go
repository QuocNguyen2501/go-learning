package service

type ElectricTrainEngine struct{}

func (c ElectricTrainEngine) Assemble() string {
	return "Assembling components of Electric Train Engine"
}
func (c ElectricTrainEngine) Capacity() int {
	return 40000
}
