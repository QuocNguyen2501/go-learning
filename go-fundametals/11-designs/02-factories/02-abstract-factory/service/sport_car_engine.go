package service

type SportCarEngine struct{}

func (s SportCarEngine) Assemble() string {
	return "Assembling components of Sport Car Engine"
}

func (s SportCarEngine) Capacity() int {
	return 10000
}

func (s SportCarEngine) Model() string {
	return "Jet F1"
}
