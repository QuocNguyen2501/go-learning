package service

type SportCarEngine struct{}

func (s SportCarEngine) Assemble() string {
	return "Assembling components of Sport Car Engine"
}

func (s SportCarEngine) Capacity() int {
	return 10000
}
